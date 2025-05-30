package kuznechikgo_test

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func TestKeySchedule(t *testing.T) {
	tt := []struct {
		masterKey string
		roundKeys [10]string
	}{
		{
			"8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef",
			[10]string{
				"8899aabbccddeeff0011223344556677", // k1
				"fedcba98765432100123456789abcdef", // k2
				"db31485315694343228d6aef8cc78c44", // k3
				"3d4553d8e9cfec6815ebadc40a9ffd04", // k4
				"57646468c44a5e28d3e59246f429f1ac", // k5
				"bd079435165c6432b532e82834da581b", // k6
				"51e640757e8745de705727265a0098b1", // k7
				"5a7925017b9fdd3ed72a91a22286f984", // k8
				"bb44e25378c73123a5f32f73cdb6e517", // k9
				"72e9dd7416bcf45b755dbaa88e4a4043", // k10
			},
		},
	}
	for _, td := range tt {
		t.Run(
			fmt.Sprintf("%s -> %s", td.masterKey, td.roundKeys),
			func(t *testing.T) {
				key, err := hex.DecodeString(td.masterKey)
				if err != nil {
					t.Fatalf("Error during decoding: %s", err)
				}

				roundKeys, err := kuznechikgo.Schedule(key)
				if err != nil {
					t.Fatalf("Error: %s", err)
				}
				var got [10]string
				for i, val := range roundKeys {
					got[i] = fmt.Sprintf("%x", val)
				}
				if !reflect.DeepEqual(got, td.roundKeys) {
					t.Errorf("\n Master key: %s\ngot:  %s\nwant: %s", td.masterKey, got, td.roundKeys)
				}
			},
		)
	}
}
