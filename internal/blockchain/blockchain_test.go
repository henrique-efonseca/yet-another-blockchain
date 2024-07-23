package blockchain

import (
	"os"
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	os.RemoveAll("testdb")
	bc := NewBlockchain("testdb")
	defer bc.Close()
	if bc == nil {
		t.Error("Blockchain should not be nil")
	}
}

func TestAddBlock(t *testing.T) {
	os.RemoveAll("testdb")
	bc := NewBlockchain("testdb")
	defer bc.Close()
	err := bc.AddBlock("First Block", nil)
	if err != nil {
		t.Error("Error adding first block:", err)
	}
	err = bc.AddBlock("Second Block", nil)
	if err != nil {
		t.Error("Error adding second block:", err)
	}
}

func TestGetLastBlock(t *testing.T) {
	os.RemoveAll("testdb")
	bc := NewBlockchain("testdb")
	defer bc.Close()
	bc.AddBlock("First Block", nil)
	lastBlock, err := bc.getLastBlock()
	if err != nil {
		t.Error("Error getting last block:", err)
	}
	if lastBlock.Data != "First Block" {
		t.Error("Last block data should be 'First Block'")
	}
}
