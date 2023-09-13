package util

import (
	"b-1-go/block"
)

const utxoBucket = "chainstate"

type UTXOSet struct {
	Blockchain *block.Blockchain
}

// FindSpendableOutputs finds and returns unspent outputs to reference in inputs
func (u UTXOSet) FindSpendableOutputs(publicKeyHash []byte, amount int) (int, map[string][]int) {
	// TODO

	return 0, nil // TODO
}
