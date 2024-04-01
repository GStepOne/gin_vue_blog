package random

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456790")

func Code(length int) string {
	rand.Seed(time.Now().UnixNano()) //指定种子，不然每次都一样
	return fmt.Sprintf("%4v", rand.Intn(10000))
}

func RandStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
