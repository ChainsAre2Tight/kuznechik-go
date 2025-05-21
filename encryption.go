package kuznechikgo

import (
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
)

func Encrypt(block []byte, keys RoundKeys) ([]byte, error) {
	fail := func(err error) ([]byte, error) {
		return nil, fmt.Errorf("block.Encrypt: %s", err)
	}

	if len(block) != 16 {
		return fail(fmt.Errorf("expected byte slice of length 16, got: %d", len(block)))
	}

	result := make([]byte, 16)
	copy(result, block)

	for i := 0; i < 9; i++ {
		transforms.ByteX(result, keys[i])
		transforms.S(result)
		transforms.L(result)
	}
	transforms.ByteX(result, keys[9])

	return result, nil
}
