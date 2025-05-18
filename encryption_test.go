package kuznechikgo_test

import (
	"fmt"
	"testing"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/utils"
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
	for _, tt := range tt {
		t.Run(
			fmt.Sprintf("%s | %s -> %s", tt.masterKey, tt.plaintext, tt.ciphertext),
			func(t *testing.T) {
				keys, err := kuznechikgo.Schedule(utils.StringToBytes(tt.masterKey))
				if err != nil {
					t.Fatalf("Error during keyschedule: %s", err)
				}
				encrypted := utils.StringToBytes(tt.plaintext)

				kuznechikgo.Encrypt(encrypted, keys)

				got := utils.BytesToString(encrypted)
				if got != tt.ciphertext {
					t.Errorf("\ngot  %s\nwant %s", got, tt.ciphertext)
				}
			},
		)
	}
}
