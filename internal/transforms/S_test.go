package transforms

import (
	"fmt"
	"testing"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/utils"
)

func TestSbox(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{"ffeeddccbbaa99881122334455667700", "b66cd8887d38e8d77765aeea0c9a7efc"},
		{"b66cd8887d38e8d77765aeea0c9a7efc", "559d8dd7bd06cbfe7e7b262523280d39"},
		{"559d8dd7bd06cbfe7e7b262523280d39", "0c3322fed531e4630d80ef5c5a81c50b"},
		{"0c3322fed531e4630d80ef5c5a81c50b", "23ae65633f842d29c5df529c13f5acda"},
	}
	for _, tt := range tt {
		t.Run(
			fmt.Sprintf("%s -> %s", tt.in, tt.out),
			func(t *testing.T) {
				res := utils.StringToBytes(tt.in)

				S(res)

				a := utils.BytesToString(res)
				if a != tt.out {
					t.Errorf("\ngot  %s\nwant %s", a, tt.out)
				}
			},
		)
	}
}

func TestInverseSbox(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{"b66cd8887d38e8d77765aeea0c9a7efc", "ffeeddccbbaa99881122334455667700"},
		{"559d8dd7bd06cbfe7e7b262523280d39", "b66cd8887d38e8d77765aeea0c9a7efc"},
		{"0c3322fed531e4630d80ef5c5a81c50b", "559d8dd7bd06cbfe7e7b262523280d39"},
		{"23ae65633f842d29c5df529c13f5acda", "0c3322fed531e4630d80ef5c5a81c50b"},
	}
	for _, tt := range tt {
		t.Run(
			fmt.Sprintf("%s -> %s", tt.in, tt.out),
			func(t *testing.T) {
				res := utils.StringToBytes(tt.in)

				InverseS(res)

				a := utils.BytesToString(res)
				if a != tt.out {
					t.Errorf("\ngot  %s\nwant %s", a, tt.out)
				}
			},
		)
	}
}
