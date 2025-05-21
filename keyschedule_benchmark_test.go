package kuznechikgo_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func BenchmarkKeyschedule(b *testing.B) {
	key, err := hex.DecodeString("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef")
	if err != nil {
		b.Fatalf("error during key decoding: %s", err)
	}
	var keys kuznechikgo.RoundKeys
	for b.Loop() {
		keys, err = kuznechikgo.Schedule(key)
		if err != nil {
			b.Fatalf("Error during keyschedule: %s", err)
		}
	}

	fmt.Println(keys)
}
