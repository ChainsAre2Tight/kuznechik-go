package scheduling

import (
	"github.com/ChainsAre2Tight/kuznechik-go/internal/tables"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/types"
)

func ScheduleKeys(masterKey []byte) types.RoundKeys {
	if len(masterKey) != 32 {
		panic("invalid master key length")
	}
	var roundKeys types.RoundKeys

	roundKeys[0] = masterKey[0:16]  // k1
	roundKeys[1] = masterKey[16:32] // k2

	for i := 1; i < 5; i++ {
		i2 := 2 * i
		roundKeys[i2] = make([]byte, 16)
		copy(roundKeys[i2], roundKeys[i2-2])
		roundKeys[i2+1] = make([]byte, 16)
		copy(roundKeys[i2+1], roundKeys[i2-1])
		for j := range 8 {
			roundKeys[i2], roundKeys[i2+1] = transforms.F(tables.Constants[8*(i-1)+j], roundKeys[i2], roundKeys[i2+1])
		}
	}

	return roundKeys
}
