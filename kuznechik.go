package kuznechikgo

import (
	"crypto/cipher"
	"fmt"

	"github.com/ChainsAre2Tight/kuznechik-go/internal/transforms"
)

var _ cipher.Block = (*Kuznechik)(nil)

type Kuznechik struct {
	keys RoundKeys
}

func New(key []byte) *Kuznechik {
	var err error
	res := &Kuznechik{}
	res.keys, err = Schedule(key)
	if err != nil {
		panic(fmt.Errorf("kuznechikgo.New: Error during keyschedule: %s", err))
	}
	return res
}

func (k *Kuznechik) BlockSize() int {
	return 16
}

func (k *Kuznechik) Decrypt(dst []byte, src []byte) {
	if l := len(dst); l != 16 {
		panic(fmt.Errorf("expected dst byte slice of length 16, got: %d", l))
	}
	if l := len(src); l != 16 {
		panic(fmt.Errorf("expected src byte slice of length 16, got: %d", l))
	}

	copy(dst, src)

	transforms.ByteX(dst, k.keys[9])
	for i := 8; i >= 0; i-- {
		transforms.InverseL(dst)
		transforms.InverseS(dst)
		transforms.ByteX(dst, k.keys[i])
	}
}

func (k *Kuznechik) Encrypt(dst []byte, src []byte) {
	if l := len(dst); l != 16 {
		panic(fmt.Errorf("expected dst byte slice of length 16, got: %d", l))
	}
	if l := len(src); l != 16 {
		panic(fmt.Errorf("expected src byte slice of length 16, got: %d", l))
	}

	copy(dst, src)

	for i := range 9 {
		transforms.ByteX(dst, k.keys[i])
		transforms.S(dst)
		transforms.L(dst)
	}
	transforms.ByteX(dst, k.keys[9])
}
