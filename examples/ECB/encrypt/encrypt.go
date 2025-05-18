package main

import (
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
	text := "fsbghjkl;jgfdsghujihsdecfghujhytfrrtfghudfhjigydsztyuijytdrsfyuiokdrestyuiojusetyuiojfdr"
	paddedText := kuznechikgo.PlaintextToBytes(text)

	ciphertext := make([]byte, len(paddedText))

	for i := range len(paddedText) / 16 {
		temp, err := kuznechikgo.Encrypt(paddedText[16*i:16*i+16], keys)
		if err != nil {
			log.Fatalf("Error during encryption: %s", err)
		}
		copy(ciphertext[16*i:16*i+16], temp)
	}

	fmt.Printf("Ciphertext: %x\n", ciphertext)
}
