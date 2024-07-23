package types

// Define common types here, e.g., Transaction
type Block struct {
	Index      int
	Timestamp  int64
	Data       interface{}
	Hash       string
	PrevHash   string
	Nonce      int
	Extensions map[string]interface{}
}

type Transaction struct {
	From   string
	To     string
	Amount int
}

type Config struct {
	Network struct {
		Port  int
		Peers []string
	}
}
