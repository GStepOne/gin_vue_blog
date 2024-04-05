package main

import "fmt"

func main() {
	slice := map[string]int{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}

	for i := range slice {
		slice[i] *= 2 // 修改切片中的元素值
	}

	fmt.Println("在循环内部打印切片：", slice) // 循环结束后的值

	for _, v := range slice {
		fmt.Println(v) // 打印循环外部的值，会是循环结束后的值
	}
}
