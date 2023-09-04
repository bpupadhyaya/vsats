package block_pkg

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// SetHash method to generate a hash of the block.
// Concatenate all the data and hash the result to obtain block hash
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{})
	hash := sha256.Sum256(headers)
	block.MyBlockHash = hash[:]
}

// NewBlock function to generate new block and return that block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

// NewBeginningBlock function to create the beginning block
func NewBeginningBlock() *Block {
	return NewBlock("data -> beginning block", []byte{})
}
