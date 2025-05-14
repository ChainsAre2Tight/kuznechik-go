package transforms

func X(value, key []byte) {
	for i := range value {
		value[i] = value[i] ^ key[i]
	}
}
