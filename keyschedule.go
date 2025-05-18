package kuznechikgo

import (
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/tables"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
)

type RoundKeys [][]byte

func Schedule(masterKey []byte) (RoundKeys, error) {
	fail := func(err error) (RoundKeys, error) {
		return nil, fmt.Errorf("kuznechikgo.Schedule: %s", err)
	}

	if len(masterKey) != 32 {
		return fail(fmt.Errorf("unexexpected masterKey length: %d, expected: 32", len(masterKey)))
	}

	var roundKeys RoundKeys = make([][]byte, 10)
	var keys []byte = make([]byte, 160)
	for i := range roundKeys {
		roundKeys[i], keys = keys[:16], keys[16:]
	}

	copy(roundKeys[0], masterKey[:16]) // k1
	copy(roundKeys[1], masterKey[16:]) // k2

	// TODO: optimize unncessary copies
	for i := 1; i < 5; i++ {
		i2 := 2 * i
		copy(roundKeys[i2], roundKeys[i2-2])
		copy(roundKeys[i2+1], roundKeys[i2-1])
		for j := range 8 {
			roundKeys[i2], roundKeys[i2+1] = transforms.F(tables.Constants[8*(i-1)+j], roundKeys[i2], roundKeys[i2+1])
		}
	}

	return roundKeys, nil
}
