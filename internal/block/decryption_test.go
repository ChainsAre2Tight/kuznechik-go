package block

import (
	"fmt"
	"testing"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/scheduling"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/utils"
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
	for _, tt := range tt {
		t.Run(
			fmt.Sprintf("%s | %s -> %s", tt.masterKey, tt.plaintext, tt.ciphertext),
			func(t *testing.T) {
				keys := scheduling.ScheduleKeys(utils.StringToBytes(tt.masterKey))
				decrypted := utils.StringToBytes(tt.ciphertext)

				Decrypt(decrypted, keys)

				got := utils.BytesToString(decrypted)
				if got != tt.plaintext {
					t.Errorf("\ngot  %s\nwant %s", got, tt.plaintext)
				}
			},
		)
	}
}
