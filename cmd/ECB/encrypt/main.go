package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/block"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/scheduling"
	"github.com/ChainsAre2Tight/kuznechik-go/internal/utils"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal(fmt.Errorf("unexpected number of arguments, expected 2, got %d", len(os.Args)-1))
	}
	key := os.Args[1]
	if len(key) != 64 {
		log.Fatal(fmt.Errorf("unexpected key length, expected 64, got %d", len(key)))
	}
	input := os.Args[2]
	// var sb strings.Builder
	// for range 100000 {
	// 	sb.Write([]byte(input))
	// }
	roundKeys := scheduling.ScheduleKeys(utils.StringToBytes(key))
	plaintext := utils.PlaintextToBytes(input)
	// plaintext := utils.PlaintextToBytes(sb.String())
	var wg sync.WaitGroup
	for i := range len(plaintext) / 16 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			block.Encrypt(plaintext[16*i:16*(i+1)], roundKeys)
		}(i)
	}
	fmt.Println(utils.BytesToString(plaintext))
}
