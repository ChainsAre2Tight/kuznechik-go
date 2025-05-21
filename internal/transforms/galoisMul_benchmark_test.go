package transforms

import (
	"fmt"
	"testing"
)

func BenchmarkGaloisMul(b *testing.B) {
	var v1, v2 byte = 10, 20
	for b.Loop() {
		v1, v2 = galoisMul(v1, v2), v1
	}
	fmt.Println(v1, v2)
}
