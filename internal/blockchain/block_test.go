package blockchain

import (
	"testing"
	"time"
)

func TestCalculateHash(t *testing.T) {
	block := &Block{
		Index:        1,
		Timestamp:    time.Now().Unix(),
		Data:         "Test Block",
		PreviousHash: "0",
		Nonce:        0,
	}
	hash := block.CalculateHash()
	if len(hash) == 0 {
		t.Error("Hash should not be empty")
	}
}

func TestNewBlock(t *testing.T) {
	block := NewBlock(1, "Test Block", "0", nil)
	if block.Index != 1 {
		t.Error("Block index should be 1")
	}
	if block.Data != "Test Block" {
		t.Error("Block data should be 'Test Block'")
	}
	if block.PreviousHash != "0" {
		t.Error("Block previous hash should be '0'")
	}
	if len(block.Hash) == 0 {
		t.Error("Block hash should not be empty")
	}
}

func TestMineBlock(t *testing.T) {
	block := NewBlock(1, "Test Block", "0", nil)
	block.MineBlock()
	if len(block.Hash) == 0 {
		t.Error("Mined block hash should not be empty")
	}
	if block.Nonce == 0 {
		t.Error("Nonce should be greater than 0 after mining")
	}
}
