package gen

import (
	"crypto/rand"
	"math/big"
)

func mustRandInt(max int) int {
	rnd, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}

	return int(rnd.Int64())
}
