package kuznechikgo

import (
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
)

func Decrypt(block []byte, keys RoundKeys) ([]byte, error) {
	fail := func(err error) ([]byte, error) {
		return nil, fmt.Errorf("kuznechikgo.Decrypt: %s", err)
	}

	if len(block) != 16 {
		return fail(fmt.Errorf("expected byte slice of length 16, got: %d", len(block)))
	}

	result := make([]byte, 16)
	copy(result, block)

	transforms.X(block, keys[9])
	for i := 8; i >= 0; i-- {
		transforms.InverseL(block)
		transforms.InverseS(block)
		transforms.X(block, keys[i])
	}

	return result, nil
}
