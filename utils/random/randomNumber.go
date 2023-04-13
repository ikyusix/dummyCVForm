package random

import (
	"math/rand"
	"time"
)

func RandNumber() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := generateRandomNumber(10000000, 99999999)
	return randomNumber
}

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
