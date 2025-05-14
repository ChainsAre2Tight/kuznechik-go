package block

import (
	"fmt"

	blck "github.com/ChainsAre2Tight/kuznechik-go/internal/block"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/types"
)

func Encrypt(plaintext []byte, keys *types.RoundKeys) ([]byte, error) {
	fail := func(err error) ([]byte, error) {
		return nil, fmt.Errorf("block.Encrypt: %s", err)
	}

	if len(plaintext) != 16 {
		return fail(fmt.Errorf("expected byte slice of length 16, got: %d", len(plaintext)))
	}

	result := make([]byte, 16)
	copy(result, plaintext)
	blck.Encrypt(plaintext, *keys)
	return result, nil
}

func Decrypt(ciphertext []byte, keys *types.RoundKeys) ([]byte, error) {
	fail := func(err error) ([]byte, error) {
		return nil, fmt.Errorf("block.Decrypt: %s", err)
	}

	if len(ciphertext) != 16 {
		return fail(fmt.Errorf("expected byte slice of length 16, got: %d", len(ciphertext)))
	}

	result := make([]byte, 16)
	copy(result, ciphertext)
	blck.Decrypt(ciphertext, *keys)
	return result, nil
}
