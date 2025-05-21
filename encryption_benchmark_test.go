package kuznechikgo_test

import (
	"encoding/hex"
	"testing"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func BenchmarkEncryption(b *testing.B) {
	key, err := hex.DecodeString("1234567890123456789012345678901234567890123456789012345678901234")
	if err != nil {
		b.Fatalf("error during key decoding: %s", err)
	}
	keys, err := kuznechikgo.Schedule(key)
	if err != nil {
		b.Fatalf("error during keyschedule: %s", err)
	}
	message, err := hex.DecodeString("12345678901234567890123456789012")
	if err != nil {
		b.Fatalf("error during message decoding: %s", err)
	}

	var res []byte
	for b.Loop() {
		res, err = kuznechikgo.Encrypt(message, keys)
		if err != nil {
			b.Fatalf("error during encryption: %s", err)
		}
	}

	if len(res) == 0 {
		b.Fatalf("len res is 0")
	}
}
