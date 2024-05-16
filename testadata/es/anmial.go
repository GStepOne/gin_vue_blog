package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := &sync.WaitGroup{}

	catCounter := uint64(0)
	dogCounter := uint64(0)
	chickenCounter := uint64(0)

	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	chickenChan := make(chan struct{}, 1)

	wg.Add(3) // 增加等待计数器

	go cat(wg, catCounter, catChan, dogChan)
	go dog(wg, dogCounter, dogChan, chickenChan)
	go chicken(wg, chickenCounter, chickenChan, catChan)

	// 先发送信号给第一个goroutine
	catChan <- struct{}{}

	wg.Wait() // 等待所有goroutine执行完毕
}

// cat dog chicken
func cat(wg *sync.WaitGroup, counter uint64, catChan <-chan struct{}, dogChan chan<- struct{}) {
	for {
		if counter > 100 {
			wg.Done()
			return
		}
		//等待小鸡渠道的信号
		<-catChan
		//收到信号之后打印小鸡
		fmt.Println("cat")
		//把counter 加1
		atomic.AddUint64(&counter, 1)
		//给猫channel发信号，该输出猫了
		dogChan <- struct{}{}
	}

}

func dog(wg *sync.WaitGroup, counter uint64, dogChan <-chan struct{}, chickenChan chan<- struct{}) {
	for {
		if counter > 100 {
			wg.Done()
			return
		}
		//等待小鸡渠道的信号
		<-dogChan
		//收到信号之后打印小鸡
		fmt.Println("dog")
		//把counter 加1
		atomic.AddUint64(&counter, 1)
		//给猫channel发信号，该输出猫了
		chickenChan <- struct{}{}
	}

}

func chicken(wg *sync.WaitGroup, counter uint64, chickenChan <-chan struct{}, catChan chan<- struct{}) {
	for {
		if counter > 100 {
			wg.Done()
			return
		}
		//等待小鸡渠道的信号
		<-chickenChan
		//收到信号之后打印小鸡
		fmt.Println("chicken")
		//把counter 加1
		atomic.AddUint64(&counter, 1)
		//给猫channel发信号，该输出猫了
		catChan <- struct{}{}
	}
}
