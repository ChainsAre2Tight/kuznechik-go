package transforms

import (
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/tables"
)

func S(value []byte) {
	for i := range value {
		value[i] = tables.DirectSbox[value[i]]
	}
}

func InverseS(value []byte) {
	for i := range value {
		value[i] = tables.InverseSbox[value[i]]
	}
}

func UintS(dst []uint64) {
	if len(dst) != 2 {
		panic(fmt.Errorf("transforms.UintS: unexpected dst length. Expected: 2, Got: %d", len(dst)))
	}

	for i := range 2 {
		var temp uint64
		for byteIndex := 0; byteIndex < 8; byteIndex++ {
			shift := (7 - byteIndex) * 8
			b := byte(dst[i] >> shift)
			sub := tables.DirectSbox[b]
			temp |= uint64(sub) << shift
		}
		dst[i] = temp
	}
}
