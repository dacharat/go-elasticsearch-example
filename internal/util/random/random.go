package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Random[T any](arr []T) T {
	return arr[rand.Intn(len(arr))]
}
