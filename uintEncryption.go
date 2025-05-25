package kuznechikgo

import (
	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
)

func UintEncrypt(upper, lower uint64, keys UintRoundKeys) (uint64, uint64) {
	result := []uint64{upper, lower}

	for i := 0; i < 9; i++ {
		transforms.UintX(result, keys[i])
		transforms.UintS(result)
		transforms.UintL(result)
	}
	transforms.UintX(result, keys[9])

	return result[0], result[1]
}
