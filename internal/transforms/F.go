package transforms

func F(key [16]byte, first, second []byte) ([]byte, []byte) {

	// TODO: find a way to avoid using temp
	temp := make([]byte, 16)
	copy(temp, first)

	ByteX(first, key[:])
	S(first)
	L(first)
	ByteX(first, second)

	copy(second, temp)
	return first, second
}
