package block_pkg

// AddBlock method to add a new block to blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	previousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, previousBlock.MyBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

// NewBlockchain function to return the whole blockchain and add the beginning block to it
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewBeginningBlock()}}
}
