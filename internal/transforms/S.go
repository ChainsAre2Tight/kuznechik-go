package transforms

import "github.com/ChainsAre2Tight/kuznechik-go/internal/tables"

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
