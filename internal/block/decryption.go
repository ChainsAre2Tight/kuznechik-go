package block

import (
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/types"
)

func Decrypt(block []byte, keys types.RoundKeys) {
	if len(block) != 16 {
		panic(fmt.Sprintf("bad block size, expected 16 got %d", len(block)))
	}
	transforms.X(block, keys[9])
	for i := 8; i >= 0; i-- {
		transforms.InverseL(block)
		transforms.InverseS(block)
		transforms.X(block, keys[i])
	}
}
