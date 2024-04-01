package random

import (
	"fmt"
	"math/rand"
	"time"
)

var stringCode = ""

func Code(length int) string {
	rand.Seed(time.Now().UnixNano()) //指定种子，不然每次都一样
	return fmt.Sprintf("%4v", rand.Intn(10000))
}

func charCode() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子为当前时间的纳秒数

	// 生成一个随机的小写字母
	lowercase := 'a' + rune(rand.Intn('z'-'a'+1))
	fmt.Printf("Random lowercase letter: %c\n", lowercase)

	// 生成一个随机的大写字母
	uppercase := 'A' + rune(rand.Intn('Z'-'A'+1))
	fmt.Printf("Random uppercase letter: %c\n", uppercase)
}
