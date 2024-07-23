package network

import (
	"encoding/json"
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/blockchain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestP2PSync(t *testing.T) {
	bc1 := blockchain.NewBlockchain("testdb1")
	defer bc1.Close()

	bc2 := blockchain.NewBlockchain("testdb2")
	defer bc2.Close()

	bc1.AddBlock("Block1", nil)

	p2pNetwork1 := NewP2PNetwork()
	p2pNetwork2 := NewP2PNetwork()

	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		blocks, _ := json.Marshal(bc1.GetAllBlocks())
		w.Write(blocks)
	}))
	defer server1.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		blocks, _ := json.Marshal(bc2.GetAllBlocks())
		w.Write(blocks)
	}))
	defer server2.Close()

	p2pNetwork1.AddNode(server2.URL)
	p2pNetwork2.AddNode(server1.URL)

	p2pNetwork1.SyncBlockchain(bc1.GetAllBlocks())
	p2pNetwork2.SyncBlockchain(bc2.GetAllBlocks())

	// Test if bc2 synced with bc1
	blocks2 := bc2.GetAllBlocks()
	if len(blocks2) != 2 || blocks2[1].Data != "Block1" {
		t.Error("Blockchain sync failed")
	}
}
