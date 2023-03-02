package utils

import (
	"hash/maphash"
	"math/rand"
)

func RandomNumber() int64 {
	r := rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	return int64(r.Int())
}
