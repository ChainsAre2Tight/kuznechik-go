package transforms

func ByteX(value, key []byte) {
	for i := range value {
		value[i] = value[i] ^ key[i]
	}
}

// TODO: vectorize
func UintX(value, key []uint64) {
	for i := range value {
		value[i] = value[i] ^ key[i]
	}
}
