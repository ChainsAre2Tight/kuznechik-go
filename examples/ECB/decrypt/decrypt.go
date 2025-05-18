package main

import (
	"encoding/hex"
	"fmt"
	"log"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func main() {
	key := []byte("12345678901234567890123456789012")
	keys, err := kuznechikgo.Schedule(key)
	if err != nil {
		log.Fatalf("Error during keyschedule: %s", err)
	}
	ciphertext, err := hex.DecodeString("66736267686a6b6c3b6a676664736768756a696873646563666768756a68797466727274666768756466686a69677964737a747975696a7974647273667975696f6b64726573747975696f6a757365747975696f6a666472010000000000000000")
	if err != nil {
		log.Fatalf("Error during ciphertext decoding: %s", err)
	}

	plaintext := make([]byte, len(ciphertext))

	for i := range len(ciphertext) / 16 {
		temp, err := kuznechikgo.Decrypt(ciphertext[16*i:16*i+16], keys)
		if err != nil {
			log.Fatalf("Error during decryption: %s", err)
		}
		copy(plaintext[16*i:16*i+16], temp)
	}

	fmt.Printf("Ciphertext: %s\n", string(plaintext))
}
