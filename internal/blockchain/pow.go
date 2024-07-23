package blockchain

import (
	"fmt"
	"strings"
)

const Difficulty = 3 // Adjust difficulty as needed

func (b *Block) MineBlock() {
	target := strings.Repeat("0", Difficulty)
	for !strings.HasPrefix(b.Hash, target) {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
	fmt.Printf("Block mined: %s\n", b.Hash)
}
