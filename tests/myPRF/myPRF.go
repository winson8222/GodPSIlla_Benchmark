package myPRF

import (
	"math/rand"
	"tests/constants"
)

// Given a rng source of random numbers and a length n, generates a str of length n
func RandStringBytesMask(n int, rng *rand.Rand) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		idx := rng.Intn(len(constants.LetterBytes))
		b[i] = constants.LetterBytes[idx]
	}
	return string(b)
}

//func make
