package transforms

func L(in []byte) {
	for range 16 {
		R(in)
	}
}

func InverseL(in []byte) {
	for range 16 {
		InverseR(in)
	}
}

func UintL(in []uint64) {
	for range 16 {
		UintR(in)
	}
}
