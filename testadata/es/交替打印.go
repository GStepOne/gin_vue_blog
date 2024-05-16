package main

import (
	"fmt"
	"sync"
)

func printOdd(wg *sync.WaitGroup, oddCh, evenCh chan int) {
	defer wg.Done()
	for i := 1; i <= 100; i += 2 {
		fmt.Println("Odd:", i)
		oddCh <- i // 将奇数发送到奇数通道
		<-evenCh   // 等待从偶数通道接收数据
	}
}

func printEven(wg *sync.WaitGroup, oddCh, evenCh chan int) {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		<-oddCh // 等待从奇数通道接收数据
		fmt.Println("Even:", i)
		evenCh <- i + 1 // 将偶数发送到偶数通道
	}
}

func main() {
	var wg sync.WaitGroup
	oddCh := make(chan int)
	evenCh := make(chan int)

	wg.Add(2)
	go printOdd(&wg, oddCh, evenCh)
	go printEven(&wg, oddCh, evenCh)

	wg.Wait()
	close(oddCh)
	close(evenCh)
}
