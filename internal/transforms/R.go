package transforms

import "github.com/ChainsAre2Tight/kuznechik-go/internal/tables"

var l = [16]byte{148, 32, 133, 16, 194, 192, 1, 251, 1, 192, 194, 16, 133, 32, 148, 1}

func linear(in []byte) byte {
	var res byte
	for i := range in {
		// fmt.Println(res)
		res ^= tables.LinearLookup[in[i]][l[i]]
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
