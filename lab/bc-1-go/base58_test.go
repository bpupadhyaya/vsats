package main

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestBase58(t *testing.T) {
	rawHash := "e58c5c56363315756ec02288eb586bdae67747dd1af0326e"
	hash, err := hex.DecodeString(rawHash)
	if err != nil {
		log.Fatal(err)
	}

	encoded := Base58Encode(hash)
	assert.Equal(t, "MvjtMw3QBmHfoKvUBqPfQ7csiR1LLwU49", string(encoded))

	decoded := Base58Decode([]byte("MvjtMw3QBmHfoKvUBqPfQ7csiR1LLwU49"))
	assert.Equal(t, strings.ToLower("e58c5c56363315756ec02288eb586bdae67747dd1af0326e"), hex.EncodeToString(decoded))

}
