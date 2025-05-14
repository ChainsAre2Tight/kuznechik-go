package keyschedule

import (
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/scheduling"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/types"
)

func Schedule(key string) (*types.RoundKeys, error) {
	fail := func(err error) (*types.RoundKeys, error) {
		return nil, fmt.Errorf("keyschedule.Schedule: %s", err)
	}

	if len(key) != 32 {
		return fail(fmt.Errorf("unexexpected key string length: %d, expected: 32", len(key)))
	}

	bytes := []byte(key)
	if len(bytes) != 32 {
		return fail(fmt.Errorf("forbidden characters found in key"))
	}

	keys := scheduling.ScheduleKeys(bytes)
	return &keys, nil
}
