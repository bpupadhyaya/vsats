package wallet

import (
	"b-1-go/util"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

const version = byte(0x00)
const addressChecksumLen = 4

// Wallet stores private and public keys
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// NewWallet function creates and returns a Wallet
func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

// GetAddress returns wallet address
func (w Wallet) GetAddress() []byte {
	publicKeyHash := HashPublicKey(w.PublicKey)
	versionedPayload := append([]byte{version}, publicKeyHash...)
	checkSum := CheckSum(versionedPayload)
	fullPayload := append(versionedPayload, checkSum...)
	address := util.Base58Encode(fullPayload)

	return address
}

// HashPublicKey hashes public key
func HashPublicKey(publicKey []byte) []byte {
	// TODO

	return nil // TODO
}

// ValidateAddress checks if address is valid
func ValidateAddress(address string) bool {
	// TODO

	return false // TODO
}

func CheckSum(payload []byte) []byte {
	// TODO

	return nil // TODO
}

// newKeyPair generates private-public key pair
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P521()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return *privateKey, publicKey
}
