package transforms

func F(key [16]byte, first, second []byte) ([]byte, []byte) {

	// TODO: find a way to avoid using temp
	temp := make([]byte, 16)
	copy(temp, first)

	X(first, key[:])
	S(first)
	L(first)
	X(first, second)
	return first, temp
}
