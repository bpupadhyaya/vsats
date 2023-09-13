package block

import (
	"b-1-go/transaction"
	"crypto/ecdsa"
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain_%s.db"
const blockBucket = "blocks"
const genesisCoinbaseData = "Dummy coinbase data"

// Blockchain implements interactions with a DB
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// TODO

// SignTransaction signs inputs of a Transaction
func (bc *Blockchain) SignTransaction(tx *transaction.Transaction, privateKey ecdsa.PrivateKey) {
	// TODO
}
