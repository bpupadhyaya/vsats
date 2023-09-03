package block_pkg

// Function to add a new block to blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	previousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, previousBlock.MyBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

// Function to return the whole blockchain and add the beginning block to it.newBlock
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewBeginningBlock()}}
}
