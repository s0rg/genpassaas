package gen

import (
	"crypto/rand"
	"math/big"
)

type Fn func(int) string

func mustRandInt(max int) int {
	rnd, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}

	return int(rnd.Int64())
}

type swapper interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s swapper) {
	for l := s.Len() - 1; l > 1; l-- {
		s.Swap(mustRandInt(l), l)
	}
	s.Swap(0, mustRandInt(s.Len()))
}
