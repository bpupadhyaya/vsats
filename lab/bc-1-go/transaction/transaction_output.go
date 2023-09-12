package transaction

import (
	"b-1-go/util"
	"bytes"
	"encoding/gob"
	"log"
)

// TxOutput represents a transaction output
type TxOutput struct {
	Value         int
	PublicKeyHash []byte
}

// Lock signs the output
func (out *TxOutput) Lock(address []byte) {
	publicKeyHash := util.Base58Decode(address)
	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-4]
	out.PublicKeyHash = publicKeyHash
}

// IsLockedWithKey checks if the output can be used by the owner of the publicKey
func (out *TxOutput) IsLockedWithKey(publicKeyHash []byte) bool {
	return bytes.Compare(out.PublicKeyHash, publicKeyHash) == 0
}

// NewTxOutput creates a new TxOutput
func NewTxOutput(value int, address string) *TxOutput {
	txOutput := &TxOutput{value, nil}
	txOutput.Lock([]byte(address))

	return txOutput
}

// TxOutputs is one or more TxOutput
type TxOutputs struct {
	Outputs []TxOutput
}

// Serialize serializes TxOutputs
func (outs TxOutputs) Serialize() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// DeserializeOutputs deserializes TxOutputs
func DeserializeOutputs(data []byte) TxOutputs {
	var outputs TxOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
