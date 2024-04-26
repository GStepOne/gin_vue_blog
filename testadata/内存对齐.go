package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int
	var y float64
	var z struct {
		a int
		b float64
		c string
	}

	fmt.Println("Alignment of int:", unsafe.Alignof(x))
	fmt.Println("Alignment of float64:", unsafe.Alignof(y))
	fmt.Println("Alignment of struct:", unsafe.Alignof(z))
}
