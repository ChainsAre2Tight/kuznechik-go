package transforms

import "github.com/ChainsAre2Tight/kuznechik-go/internal/tables"

func galoisMul(a, b uint8) uint8 {
	// fmt.Println(a, b, GaloisInv[a], GaloisInv[b], (GaloisInv[a] + GaloisInv[b]), Galois[GaloisInv[a]+GaloisInv[b]])
	if a == 0 || b == 0 {
		return 0
	}
	var res uint8

	// account for integer overflow
	dif := 255 - tables.GaloisFieldLogarithm[a]
	if tables.GaloisFieldLogarithm[b] > dif {
		res = tables.GaloisFieldExponent[tables.GaloisFieldLogarithm[a]+tables.GaloisFieldLogarithm[b]+1]
	} else {
		res = tables.GaloisFieldExponent[tables.GaloisFieldLogarithm[a]+tables.GaloisFieldLogarithm[b]]
	}

	return res
}

var l = [16]byte{148, 32, 133, 16, 194, 192, 1, 251, 1, 192, 194, 16, 133, 32, 148, 1}

func linear(in []byte) byte {
	var res byte
	for i := range in {
		// fmt.Println(res)
		res ^= galoisMul(in[i], l[i])
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
