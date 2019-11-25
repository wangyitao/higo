package util

import (
	trand "crypto/rand"
	"math/big"
	prand "math/rand"
	"strconv"
	"time"
)

// PseudoRandStringRunes 返回随机字符串，伪随机,在1s内产生的字符串是一样的
func PseudoRandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	prand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[prand.Intn(len(letterRunes))]
	}
	return string(b)
}

// TrueRandStringRunes 返回随机字符串，真随机，每次产生的字符串都不一样
func TrueRandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for j := range b {
		result, _ := trand.Int(trand.Reader, big.NewInt(int64(len(letterRunes))))
		i, err := strconv.Atoi(result.String())
		if err == nil {
			b[j] = letterRunes[i]
		}
	}
	return string(b)
}
