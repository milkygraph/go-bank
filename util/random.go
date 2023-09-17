package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[RandInt(0, len(alphabet))]
	}
	return string(b)
}

func RandOwner() string {
	return RandString(6)
}

func RandMoney() int64 {
	return int64(RandInt(0, 1000))
}

func RandCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[RandInt(0, n)]
}

func RandEmail() string {
	return RandString(6) + "@email.com"
}
