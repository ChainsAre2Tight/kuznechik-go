package kuznechikgo_test

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"testing"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func TestEndToEnd(t *testing.T) {
	tt := []struct {
		masterKey  string
		plaintext  string
		ciphertext string
	}{
		{
			"8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef",
			"1122334455667700ffeeddccbbaa9988",
			"7f679d90bebc24305a468d42b9d4edcd",
		},
	}
	for _, td := range tt {
		t.Run(
			fmt.Sprintf("%s | %s -> %s", td.masterKey, td.plaintext, td.ciphertext),
			func(t *testing.T) {
				mk, err := hex.DecodeString(td.masterKey)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}
				plaintext, err := hex.DecodeString(td.plaintext)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}
				ciphertext, err := hex.DecodeString(td.ciphertext)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}

				kuzya := kuznechikgo.New(mk)
				res := make([]byte, 16)
				kuzya.Encrypt(res, plaintext)

				if !bytes.Equal(res, ciphertext) {
					t.Errorf("\nencryption: \ngot  %s\nwant %s", res, td.ciphertext)
				}

				kuzya.Decrypt(res, ciphertext)

				if !bytes.Equal(res, plaintext) {
					t.Errorf("\ndecryption: \ngot  %s\nwant %s", res, td.ciphertext)
				}
			},
		)
	}
}

func BenchmarkEncrypt(b *testing.B) {
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, key)
	blk := make([]byte, 16)

	c := kuznechikgo.New(key)

	for b.Loop() {
		c.Encrypt(blk, blk)
	}
}
