package transforms

import (
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/tables"
)

var l = [16]byte{148, 32, 133, 16, 194, 192, 1, 251, 1, 192, 194, 16, 133, 32, 148, 1}

func linear(in []byte) byte {
	var res byte
	for i := range in {
		// fmt.Println(res)
		res ^= tables.LinearLookup[in[i]][l[i]]
	}
	return res
}

func uintLinear(in []uint64) uint64 {
	var res uint64

	upper := in[0]
	lower := in[1]

	first := 7
	second := 15
	for range 8 {
		res ^= uint64(tables.LinearLookup[byte(lower)][l[second]])
		res ^= uint64(tables.LinearLookup[byte(upper)][l[first]])

		upper >>= 8
		lower >>= 8

		first--
		second--
	}
	return res
}

func R(in []byte) {
	temp := linear(in)

	// left byteshift
	for i := 15; i > 0; i-- {
		in[i] = in[i-1]
	}
	in[0] = temp
}

func InverseR(in []byte) {
	// right byteshift
	temp := in[0]
	for i := range 15 {
		in[i] = in[i+1]
	}
	in[15] = temp

	temp = linear(in)
	in[15] = temp
}

func UintR(in []uint64) {
	if len(in) != 2 {
		panic(fmt.Errorf("transforms.UintR: unexpected dst length. Expected: 2, Got: %d", len(in)))
	}

	first := uintLinear(in) << 56

	// right byteshift
	in[1] >>= 8
	carry := uint64(byte(in[0])) << 56
	in[1] |= carry
	in[0] >>= 8

	in[0] |= first
}
