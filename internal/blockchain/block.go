package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Block struct {
	Index        int                    `json:"index"`
	Timestamp    int64                  `json:"timestamp"`
	Data         interface{}            `json:"data"` // Use interface{} for flexible data
	PreviousHash string                 `json:"previousHash"`
	Hash         string                 `json:"hash"`
	Nonce        int                    `json:"nonce"`
	Extensions   map[string]interface{} `json:"extensions"` // Extensions for custom fields
}

func (b *Block) CalculateHash() string {
	record := string(b.Index) + string(b.Timestamp) + b.PreviousHash + string(b.Nonce)
	if b.Data != nil {
		dataBytes, _ := json.Marshal(b.Data)
		record += string(dataBytes)
	}
	if b.Extensions != nil {
		extBytes, _ := json.Marshal(b.Extensions)
		record += string(extBytes)
	}
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewBlock(index int, data interface{}, previousHash string, extensions map[string]interface{}) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
		Nonce:        0,
		Extensions:   extensions,
	}
	block.Hash = block.CalculateHash()
	return block
}
