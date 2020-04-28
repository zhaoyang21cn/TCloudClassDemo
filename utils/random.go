package utils

import (
	"math/rand"
	"time"
)

// Random 生成随机数
func Random() int {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	random := r.Intn(100000)
	return random
}