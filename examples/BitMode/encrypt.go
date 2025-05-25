package main

import (
	"fmt"
	"log"

	kuznechikgo "github.com/ChainsAre2Tight/kuznechik-go"
)

func main() {
	var upper, lower uint64
	upper = 0x1122334455667788
	lower = 0x9900aabbccddeeff

	key := make([]byte, 32)
	keys, err := kuznechikgo.Schedule(key)
	if err != nil {
		log.Fatalf("Error during keyschedule: %s", err)
	}
	UintKeys := kuznechikgo.KeysToUints(keys)

	upper, lower = kuznechikgo.UintEncrypt(upper, lower, UintKeys)
	fmt.Printf("Encrypted: %0.16x%0.16x\n", upper, lower)
}
