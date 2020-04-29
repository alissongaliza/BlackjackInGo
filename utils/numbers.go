package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GetRandomNumber(start int, end int) (number int) {
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(end-start+1) + start
	return
}

func StringToInt(str string) int {

	if s, err := strconv.Atoi(str); err == nil {
		return s
	} else {
		panic(err)
	}
}
