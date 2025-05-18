package kuznechikgo_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func TestBlockEncryption(t *testing.T) {
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
				keys, err := kuznechikgo.Schedule(mk)
				if err != nil {
					t.Fatalf("Error during keyschedule: %s", err)
				}
				encrypted, err := hex.DecodeString(td.plaintext)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}

				kuznechikgo.Encrypt(encrypted, keys)

				got := fmt.Sprintf("%x", encrypted)
				if got != td.ciphertext {
					t.Errorf("\ngot  %s\nwant %s", got, td.ciphertext)
				}
			},
		)
	}
}
