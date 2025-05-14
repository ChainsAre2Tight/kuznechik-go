package block

import (
	"fmt"
	"testing"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/scheduling"
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
				keys := scheduling.ScheduleKeys(utils.StringToBytes(tt.masterKey))
				encrypted := utils.StringToBytes(tt.plaintext)

				Encrypt(encrypted, keys)

				got := utils.BytesToString(encrypted)
				if got != tt.ciphertext {
					t.Errorf("\ngot  %s\nwant %s", got, tt.ciphertext)
				}
			},
		)
	}
}
