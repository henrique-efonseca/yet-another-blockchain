package network

import (
	"log"
	"net"
	"sync"
)

type Peer struct {
	Address string
	Conn    net.Conn
}

type PeerManager struct {
	Peers map[string]*Peer
	mutex sync.Mutex
}

func NewPeerManager() *PeerManager {
	return &PeerManager{
		Peers: make(map[string]*Peer),
	}
}

func (pm *PeerManager) AddPeer(address string, conn net.Conn) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if _, exists := pm.Peers[address]; exists {
		return // Peer already exists
	}

	pm.Peers[address] = &Peer{
		Address: address,
		Conn:    conn,
	}
	log.Printf("Peer added: %s", address)
}

func (pm *PeerManager) RemovePeer(address string) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if peer, exists := pm.Peers[address]; exists {
		peer.Conn.Close()
		delete(pm.Peers, address)
		log.Printf("Peer removed: %s", address)
	}
}

func (pm *PeerManager) GetPeerList() []string {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	peers := make([]string, 0, len(pm.Peers))
	for address := range pm.Peers {
		peers = append(peers, address)
	}
	return peers
}
