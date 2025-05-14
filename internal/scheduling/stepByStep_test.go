package scheduling

import (
	"testing"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/tables"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/utils"
)

func TestKeyScheduleStepByStep(t *testing.T) {
	key := utils.StringToBytes("8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef")
	k1 := key[0:16]
	k2 := key[16:32]
	k3 := make([]byte, 16)
	k4 := make([]byte, 16)
	copy(k3, k1)
	copy(k4, k2)
	got1, got2 := transforms.F(tables.Constants[0], k3, k4)
	got1s := utils.BytesToString(got1)
	got2s := utils.BytesToString(got2)
	want1 := "c3d5fa01ebe36f7a9374427ad7ca8949"
	want2 := "8899aabbccddeeff0011223344556677"
	if got1s != want1 || got2s != want2 {
		t.Errorf("Step 1:\ngot  %s | %s\nwant %s | %s", got1s, got2s, want1, want2)
	}
}
