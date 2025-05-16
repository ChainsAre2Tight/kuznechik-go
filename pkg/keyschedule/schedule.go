package keyschedule

import (
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/scheduling"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/utils"
	"github.com/ChainsAre2Tight/kuznechik-go/pkg/types"
)

func Schedule(key string) (*types.RoundKeys, error) {
	fail := func(err error) (*types.RoundKeys, error) {
		return nil, fmt.Errorf("keyschedule.Schedule: %s", err)
	}

	if len(key) != 64 {
		return fail(fmt.Errorf("unexexpected key string length: %d, expected: 64", len(key)))
	}

	bytes := utils.StringToBytes(key)
	if len(bytes) != 32 {
		return fail(fmt.Errorf("forbidden characters found in key"))
	}

	keys := scheduling.ScheduleKeys(bytes)
	return &keys, nil
}
