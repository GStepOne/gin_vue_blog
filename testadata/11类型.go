package main

import (
	"fmt"
)

func main() {
	var str string = "Hello, world!"
	var num int = 42
	var flt float64 = 3.14
	var boolean bool = true
	var arr []int = []int{1, 2, 3}
	var mp map[string]int = map[string]int{"a": 1, "b": 2}
	var ch chan int = make(chan int)
	var fn func() = func() { fmt.Println("Hello from function!") }
	var ptr *int = &num
	var iface interface{} = "Hello"

	printType(str)
	printType(num)
	printType(flt)
	printType(boolean)
	printType(arr)
	printType(mp)
	printType(ch)
	printType(fn)
	printType(ptr)
	printType(iface)
}

func printType(v interface{}) {
	fmt.Printf("Type of %v is %T\n", v, v)
}
