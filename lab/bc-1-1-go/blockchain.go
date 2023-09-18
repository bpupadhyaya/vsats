package main

// Blockchain is a sequence of blocks
type Blockchain struct {
	blocks []*Block
}

// AddBlock adds supplied data as a block to the existing blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with beginning block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewBeginningBlock()}}
}
