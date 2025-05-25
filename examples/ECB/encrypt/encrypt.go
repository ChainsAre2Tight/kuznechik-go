package main

import (
	"fmt"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func main() {
	key := []byte("12345678901234567890123456789012")

	kuzya := kuznechikgo.New(key)

	text := "fsbghjkl;jgfdsghujihsdecfghujhytfrrtfghudfhjigydsztyuijytdrsfyuiokdrestyuiojusetyuiojfdr"
	paddedText := kuznechikgo.PlaintextToBytes(text)

	ciphertext := make([]byte, len(paddedText))

	for i := range len(paddedText) / 16 {
		kuzya.Encrypt(ciphertext[16*i:16*i+16], paddedText[16*i:16*i+16])
	}

	fmt.Printf("Ciphertext: %x\n", ciphertext)
}
