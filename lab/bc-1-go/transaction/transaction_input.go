package transaction

import (
	"b-1-go/wallet"
	"bytes"
)

// TxInput represents a transaction input
type TxInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PublicKey []byte
}

// UsesKey checks whether the address initiated the transaction
func (in *TxInput) UsesKey(publicKeyHash []byte) bool {
	lockingHash := wallet.HashPublicKey(in.PublicKey)
	return bytes.Compare(lockingHash, publicKeyHash) == 0
}
