package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func BytesToString(in []byte) string {
	var sb strings.Builder
	for _, value := range in {
		sb.Write([]byte(fmt.Sprintf("%02x", value)))
	}
	return sb.String()
}

func StringToBytes(in string) []byte {
	res := make([]byte, len(in)/2)
	for i := range len(in) / 2 {
		a, err := strconv.ParseUint(in[2*i:2*i+2], 16, 8)
		if err != nil {
			fmt.Println(a, err, in[i:i+2])
		}
		res[i] = uint8(a)
	}
	return res
}

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
