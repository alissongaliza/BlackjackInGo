package utils

import (
	"math/rand"
	"time"
)

func GetRandomNumber(start int, end int) (number int) {
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(end-start+1) + start
	return
}
