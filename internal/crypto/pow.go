package crypto

import (
	"strconv"
	"strings"
)

const Difficulty = 1

func ProofOfWork(data string) (string, int) {
	var nonce int
	var hash string
	for {
		nonce++
		hash = CalculateHash(data + strconv.Itoa(nonce))
		if strings.HasPrefix(hash, strings.Repeat("0", Difficulty)) {
			break
		}
	}
	return hash, nonce
}
