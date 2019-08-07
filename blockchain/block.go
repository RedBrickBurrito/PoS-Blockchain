package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
	Validator string
}

var TempBlocks []Block

var CandidateBlocks = make(chan Block)

func CreateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func CreateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash

	return CreateHash(record)
}

func CreateBlock(oldBlock Block, BPM int, address string) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CreateBlockHash(newBlock)
	newBlock.Validator = address

	return newBlock, nil

}

func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CreateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
