package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha512"
	"golang.org/x/crypto/ripemd160"
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
	address := Base58Encode(fullPayload)

	return address
}

// HashPublicKey hashes public key
func HashPublicKey(publicKey []byte) []byte {
	publicSHA512 := sha512.Sum512(publicKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA512[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

// ValidateAddress checks if address is valid
func ValidateAddress(address string) bool {
	publicKeyHash := Base58Decode([]byte(address))
	actualCheckSum := publicKeyHash[len(publicKeyHash)-addressChecksumLen:]
	version := publicKeyHash[0]
	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-addressChecksumLen]
	targetCheckSum := CheckSum(append([]byte{version}, publicKeyHash...))

	return bytes.Compare(actualCheckSum, targetCheckSum) == 0
}

func CheckSum(payload []byte) []byte {
	firstSHA := sha512.Sum512(payload)
	secondSHA := sha512.Sum512(firstSHA[:])

	return secondSHA[:addressChecksumLen]
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
