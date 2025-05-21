package transforms_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
)

func TestR(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{"00000000000000000000000000000100", "94000000000000000000000000000001"},
		{"94000000000000000000000000000001", "a5940000000000000000000000000000"},
		{"a5940000000000000000000000000000", "64a59400000000000000000000000000"},
		{"64a59400000000000000000000000000", "0d64a594000000000000000000000000"},
	}
	for _, tt := range tt {
		t.Run(
			fmt.Sprintf("%s -> %s", tt.in, tt.out),
			func(t *testing.T) {
				res, err := hex.DecodeString(tt.in)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}

				transforms.R(res)

				a := fmt.Sprintf("%x", res)
				if a != tt.out {
					t.Errorf("\ngot  %s\nwant %s", a, tt.out)
				}
			},
		)
	}
}

func TestInverseR(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{"0d64a594000000000000000000000000", "64a59400000000000000000000000000"},
		{"64a59400000000000000000000000000", "a5940000000000000000000000000000"},
		{"a5940000000000000000000000000000", "94000000000000000000000000000001"},
		{"94000000000000000000000000000001", "00000000000000000000000000000100"},
	}
	for _, tt := range tt {
		t.Run(
			fmt.Sprintf("%s -> %s", tt.in, tt.out),
			func(t *testing.T) {
				res, err := hex.DecodeString(tt.in)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}

				transforms.InverseR(res)

				a := fmt.Sprintf("%x", res)
				if a != tt.out {
					t.Errorf("\ngot  %s\nwant %s", a, tt.out)
				}
			},
		)
	}
}

func BenchmarkR(b *testing.B) {
	value, err := hex.DecodeString("64a59400000000000000000000000000")
	if err != nil {
		b.Fatalf("error during value decoding")
	}
	for b.Loop() {
		transforms.R(value)
	}
}
