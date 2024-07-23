package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// CalculateHash generates a SHA-256 hash for a given input string
func CalculateHash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
