package block

import (
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/types"
)

func Encrypt(block []byte, keys types.RoundKeys) {
	if len(block) != 16 {
		panic(fmt.Sprintf("bad block size, expected 16 got %d", len(block)))
	}
	for i := 0; i < 9; i++ {
		transforms.X(block, keys[i])
		transforms.S(block)
		transforms.L(block)
	}
	transforms.X(block, keys[9])
}
