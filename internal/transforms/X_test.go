package transforms_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
)

func BenchmarkX(b *testing.B) {
	block, err := hex.DecodeString("12345678901234567890123456789012")
	if err != nil {
		b.Fatalf("error during block decoding: %s", err)
	}
	key := make([]byte, 16)
	copy(key, block)
	for b.Loop() {
		transforms.X(block, key)
	}

	fmt.Println(block)
}
