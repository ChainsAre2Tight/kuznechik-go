package transforms

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestL(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{"64a59400000000000000000000000000", "d456584dd0e3e84cc3166e4b7fa2890d"},
		{"d456584dd0e3e84cc3166e4b7fa2890d", "79d26221b87b584cd42fbc4ffea5de9a"},
		{"79d26221b87b584cd42fbc4ffea5de9a", "0e93691a0cfc60408b7b68f66b513c13"},
		{"0e93691a0cfc60408b7b68f66b513c13", "e6a8094fee0aa204fd97bcb0b44b8580"},
	}
	for _, tt := range tt {
		t.Run(
			fmt.Sprintf("%s -> %s", tt.in, tt.out),
			func(t *testing.T) {
				res, err := hex.DecodeString(tt.in)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}

				L(res)

				a := fmt.Sprintf("%x", res)
				if a != tt.out {
					t.Errorf("\ngot  %s\nwant %s", a, tt.out)
				}
			},
		)
	}
}

func TestInverseL(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{"e6a8094fee0aa204fd97bcb0b44b8580", "0e93691a0cfc60408b7b68f66b513c13"},
		{"0e93691a0cfc60408b7b68f66b513c13", "79d26221b87b584cd42fbc4ffea5de9a"},
		{"79d26221b87b584cd42fbc4ffea5de9a", "d456584dd0e3e84cc3166e4b7fa2890d"},
		{"d456584dd0e3e84cc3166e4b7fa2890d", "64a59400000000000000000000000000"},
	}
	for _, tt := range tt {
		t.Run(
			fmt.Sprintf("%s -> %s", tt.in, tt.out),
			func(t *testing.T) {
				res, err := hex.DecodeString(tt.in)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}

				InverseL(res)

				a := fmt.Sprintf("%x", res)
				if a != tt.out {
					t.Errorf("\ngot  %s\nwant %s", a, tt.out)
				}
			},
		)
	}
}
