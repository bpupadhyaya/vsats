package main

import (
	"bc-0-go/lib"
	"fmt"
)

func main() {
	newBlockChain := block_pkg.NewBlockchain()
	newBlockChain.AddBlock("data -> first transaction")
	newBlockChain.AddBlock("data -> second transaction")

	for i, block := range newBlockChain.Blocks {
		fmt.Printf("Block ID: %d \n", i)
		fmt.Printf("Timestamp: %d \n", block.Timestamp+int64(i))
		fmt.Printf("Hash of the block: %x\n", block.MyBlockHash)
		fmt.Printf("Hash of the previous block: %x\n", block.PreviousBlockHash)
		fmt.Printf("Transactions: %s\n", block.AllData)
	}
}
