package transforms

import (
	"fmt"
	"testing"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/utils"
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
				res := utils.StringToBytes(tt.in)

				R(res)

				a := utils.BytesToString(res)
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
				res := utils.StringToBytes(tt.in)

				InverseR(res)

				a := utils.BytesToString(res)
				if a != tt.out {
					t.Errorf("\ngot  %s\nwant %s", a, tt.out)
				}
			},
		)
	}
}
