package main

import "fmt"

func main() {
	blockChain := NewBlockchain()
	blockChain.AddBlock("My first BTC send, 0.5BTC")
	blockChain.AddBlock("My second BTC send, 1BTC")
	var i int = 1
	for _, block := range blockChain.blocks {
		fmt.Printf("count: %d\n", i)
		fmt.Printf("Previous hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		i++
	}
}
