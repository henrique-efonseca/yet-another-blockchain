package network

import (
	"encoding/json"
	"fmt"
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/blockchain"
	"github.com/spf13/viper"
	"log"
	"net"
	"sync"
)

var GlobalNetwork *P2PNetwork

type P2PNetwork struct {
	Blockchain  *blockchain.Blockchain
	PeerManager *PeerManager
	mutex       sync.Mutex
}

func NewP2PNetwork(bc *blockchain.Blockchain) *P2PNetwork {
	return &P2PNetwork{
		Blockchain:  bc,
		PeerManager: NewPeerManager(),
	}
}

func (p2p *P2PNetwork) StartServer(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer ln.Close()

	log.Printf("P2P server listening on port %d", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		log.Println("Accepted new connection")
		go p2p.handleConnection(conn)
	}
}

func (p2p *P2PNetwork) handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("Handling new connection...")

	decoder := json.NewDecoder(conn)
	var message map[string]interface{}
	if err := decoder.Decode(&message); err != nil {
		log.Printf("Error decoding message: %v", err)
		return
	}

	log.Printf("Received message: %v", message)

	switch message["type"] {
	case "block":
		var block blockchain.Block
		blockData, _ := json.Marshal(message["data"])
		log.Printf("Block data received: %s", blockData)
		if err := json.Unmarshal(blockData, &block); err != nil {
			log.Printf("Error unmarshalling block: %v", err)
			return
		}

		log.Printf("Adding block: %v", block)
		if err := p2p.Blockchain.AddBlock(block.Data, block.Extensions); err != nil {
			log.Printf("Error adding block: %v", err)
			return
		}

	case "peerList":
		var peers []string
		peersData, _ := json.Marshal(message["data"])
		log.Printf("Peers data received: %s", peersData)
		if err := json.Unmarshal(peersData, &peers); err != nil {
			log.Printf("Error unmarshalling peers: %v", err)
			return
		}

		log.Printf("Adding peers: %v", peers)
		for _, peer := range peers {
			p2p.AddPeer(peer)
		}
	}
}

func (p2p *P2PNetwork) AddPeer(address string) error {
	p2p.mutex.Lock()
	defer p2p.mutex.Unlock()

	if _, exists := p2p.PeerManager.Peers[address]; exists {
		log.Printf("Peer %s already connected", address)
		return nil // Already connected to this peer
	}

	log.Printf("Connecting to peer %s", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("Error connecting to peer %s: %v", address, err)
		return err
	}
	log.Printf("Connected to peer %s", address)
	p2p.PeerManager.AddPeer(address, conn)

	// Send the current list of known peers to the new peer
	peers := p2p.PeerManager.GetPeerList()

	encoder := json.NewEncoder(conn)
	message := map[string]interface{}{
		"type": "peerList",
		"data": peers,
	}
	if err := encoder.Encode(&message); err != nil {
		log.Printf("Error sending peer list to %s: %v", address, err)
		return err
	}

	return nil
}

func (p2p *P2PNetwork) BroadcastBlock(block *blockchain.Block) {
	p2p.mutex.Lock()
	defer p2p.mutex.Unlock()

	log.Printf("Broadcasting block: %v", block)
	for _, peer := range p2p.PeerManager.Peers {
		encoder := json.NewEncoder(peer.Conn)
		message := map[string]interface{}{
			"type": "block",
			"data": block,
		}
		if err := encoder.Encode(&message); err != nil {
			log.Printf("Error sending block to peer: %v", err)
		}
	}
}

func (p2p *P2PNetwork) ConnectToBootstrapNodes() {
	bootstrapNodes := viper.GetStringSlice("bootstrapNodes")
	for _, node := range bootstrapNodes {
		err := p2p.AddPeer(node)
		if err != nil {
			log.Printf("Failed to connect to bootstrap node %s: %v", node, err)
		} else {
			log.Printf("Connected to bootstrap node %s", node)
		}
	}
}
