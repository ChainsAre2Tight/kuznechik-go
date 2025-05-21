package kuznechikgo_test

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func TestUintEncryption(t *testing.T) {
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
				k, err := kuznechikgo.Schedule(mk)
				if err != nil {
					t.Fatalf("Error during keyschedule: %s", err)
				}
				keys := kuznechikgo.KeysToUints(k)
				plaintext, err := hex.DecodeString(td.plaintext)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}

				upper := binary.BigEndian.Uint64(plaintext[:8])
				lower := binary.BigEndian.Uint64(plaintext[8:])

				encUpper, encLower, err := kuznechikgo.UintEncrypt(upper, lower, keys)
				if err != nil {
					t.Fatalf("Error during encryption: %s", err)
				}

				got := fmt.Sprintf("%0.16x%0.16x", encUpper, encLower)
				if got != td.ciphertext {
					t.Errorf("\ngot  %s\nwant %s", got, td.ciphertext)
				}
			},
		)
	}
}
