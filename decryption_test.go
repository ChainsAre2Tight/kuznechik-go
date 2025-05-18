package kuznechikgo_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func TestBlockDecryption(t *testing.T) {
	tt := []struct {
		masterKey  string
		ciphertext string
		plaintext  string
	}{
		{
			"8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef",
			"7f679d90bebc24305a468d42b9d4edcd",
			"1122334455667700ffeeddccbbaa9988",
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
				keys, err := kuznechikgo.Schedule(mk)
				if err != nil {
					t.Fatalf("Error during keyschedule: %s", err)
				}
				ciphertext, err := hex.DecodeString(td.ciphertext)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}

				plaintext, err := kuznechikgo.Decrypt(ciphertext, keys)
				if err != nil {
					t.Fatalf("Error during decryption: %s", err)
				}

				got := fmt.Sprintf("%x", plaintext)
				if got != td.plaintext {
					t.Errorf("\ngot  %s\nwant %s", got, td.plaintext)
				}
			},
		)
	}
}
