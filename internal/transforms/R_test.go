package transforms

import (
	"encoding/hex"
	"fmt"
	"testing"
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

				R(res)

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

				InverseR(res)

				a := fmt.Sprintf("%x", res)
				if a != tt.out {
					t.Errorf("\ngot  %s\nwant %s", a, tt.out)
				}
			},
		)
	}
}
