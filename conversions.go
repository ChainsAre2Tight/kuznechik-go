package kuznechikgo

func PlaintextToBytes(plaintext string) []byte {
	res := []byte(plaintext)

	padded := AddPadding(res, 16)

	return padded
}

func AddPadding(in []byte, blockSize int) []byte {
	res := append(in, 1)
	dif := blockSize - len(in)%blockSize
	for range dif {
		res = append(res, 0)
	}
	return res
}

func RemovePadding(in []byte) []byte {
	for i := len(in) - 1; i > 0; i-- {
		if in[i] == 1 {
			return in[0:i]
		}
	}
	return nil
}
