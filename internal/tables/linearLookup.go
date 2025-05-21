package tables

var LinearLookup = [256][256]byte{}

func init() {
	for upper := range 256 {
		for lower := range 256 {
			LinearLookup[upper][lower] = galoisMul(uint8(upper), uint8(lower))
		}
	}
}

func galoisMul(a, b uint8) uint8 {
	// fmt.Println(a, b, GaloisInv[a], GaloisInv[b], (GaloisInv[a] + GaloisInv[b]), Galois[GaloisInv[a]+GaloisInv[b]])
	if a == 0 || b == 0 {
		return 0
	}
	var res uint8

	// account for integer overflow
	logA := GaloisFieldLogarithm[a]
	logB := GaloisFieldLogarithm[b]
	sum := logA + logB

	if logB > 255-logA {
		res = GaloisFieldExponent[sum+1]
	} else {
		res = GaloisFieldExponent[sum]
	}

	return res
}
