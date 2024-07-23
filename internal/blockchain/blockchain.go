package blockchain

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

const levelDBPath = "./data/leveldb"

type Blockchain struct {
	db *leveldb.DB
}

func NewBlockchain() *Blockchain {
	db, err := leveldb.OpenFile(levelDBPath, nil)
	if err != nil {
		log.Fatalf("Failed to open LevelDB: %v", err)
	}
	bc := &Blockchain{db: db}

	// Check if the blockchain is empty and create a genesis block if necessary
	_, err = bc.getLastBlock()
	if err != nil {
		bc.createGenesisBlock()
	}

	return bc
}

func (bc *Blockchain) Close() {
	bc.db.Close()
}

func (bc *Blockchain) AddBlock(data interface{}, extensions map[string]interface{}) error {
	log.Println("Adding new block with data:", data)
	prevBlock, err := bc.getLastBlock()
	if err != nil {
		log.Println("Error getting last block:", err)
		return err
	}
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash, extensions)
	newBlock.MineBlock() // Assuming MineBlock is a method implemented in block.go
	newBlockHash, err := bc.saveBlock(newBlock)
	if err != nil {
		log.Println("Error saving block:", err)
		return err
	}
	err = bc.db.Put([]byte("lastHash"), []byte(newBlockHash), nil)
	if err != nil {
		log.Println("Error updating lastHash:", err)
		return err
	}
	log.Println("New block added with hash:", newBlockHash)
	return nil
}

func (bc *Blockchain) getLastBlock() (*Block, error) {
	lastHash, err := bc.db.Get([]byte("lastHash"), nil)
	if err != nil {
		log.Println("No last hash found, creating genesis block")
		return nil, err
	}
	return bc.getBlockByHash(string(lastHash))
}

func (bc *Blockchain) createGenesisBlock() {
	log.Println("Creating genesis block")
	genesisBlock := NewBlock(0, "Genesis Block", "0", nil)
	genesisBlock.MineBlock() // Assuming MineBlock is a method implemented in block.go
	genesisBlockHash, err := bc.saveBlock(genesisBlock)
	if err != nil {
		log.Fatalf("Failed to create genesis block: %v", err)
	}
	err = bc.db.Put([]byte("lastHash"), []byte(genesisBlockHash), nil)
	if err != nil {
		log.Fatalf("Failed to save genesis block hash: %v", err)
	}
	log.Println("Genesis block created with hash:", genesisBlockHash)
}

func (bc *Blockchain) saveBlock(block *Block) (string, error) {
	blockData, err := json.Marshal(block)
	if err != nil {
		log.Println("Error marshalling block:", err)
		return "", err
	}
	err = bc.db.Put([]byte(block.Hash), blockData, nil)
	if err != nil {
		log.Println("Error saving block to db:", err)
		return "", err
	}
	log.Println("Block saved with hash:", block.Hash)
	return block.Hash, nil
}

func (bc *Blockchain) getBlockByHash(hash string) (*Block, error) {
	log.Println("Getting block by hash:", hash)
	blockData, err := bc.db.Get([]byte(hash), nil)
	if err != nil {
		log.Println("Error getting block by hash:", err)
		return nil, err
	}
	var block Block
	err = json.Unmarshal(blockData, &block)
	if err != nil {
		log.Println("Error unmarshalling block data:", err)
		return nil, err
	}
	log.Println("Block found:", block)
	return &block, nil
}

func (bc *Blockchain) GetAllBlocks() ([]*Block, error) {
	log.Println("Getting all blocks")
	var blocks []*Block

	iter := bc.db.NewIterator(nil, nil)
	for iter.Next() {
		key := string(iter.Key())
		if key == "lastHash" {
			continue
		}
		value := string(iter.Value())
		log.Printf("Key: %s, Raw block data: %s", key, value) // Log the key and raw block data
		var block Block
		err := json.Unmarshal(iter.Value(), &block)
		if err != nil {
			log.Printf("Error unmarshalling block for key %s: %v", key, err) // Log the error
			return nil, err
		}
		blocks = append(blocks, &block)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Printf("Iterator error: %v", err) // Log the error
		return nil, err
	}

	log.Println("All blocks retrieved:", blocks)
	return blocks, nil
}

func (bc *Blockchain) IsValid() bool {
	log.Println("Validating blockchain")
	iter := bc.db.NewIterator(nil, nil)
	var prevBlock *Block

	for iter.Next() {
		key := string(iter.Key())
		if key == "lastHash" {
			continue
		}
		var block Block
		err := json.Unmarshal(iter.Value(), &block)
		if err != nil {
			log.Println("Error unmarshalling block:", err)
			return false
		}

		if prevBlock != nil {
			if block.PreviousHash != prevBlock.Hash {
				log.Println("Invalid blockchain: previous hash does not match")
				return false
			}
		}

		prevBlock = &block
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		log.Println("Iterator error:", err)
		return false
	}

	log.Println("Blockchain is valid")
	return true
}
