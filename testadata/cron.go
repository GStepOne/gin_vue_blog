package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	Cron := cron.New(cron.WithSeconds())
	//Cron.AddFunc("* * * * * *", Func1)
	//Cron.AddFunc("*/2 * * * * *", inner("jack"))
	Cron.AddJob("*/2 * * * * *", DummyJob{Name: "jack"})
	Cron.Start()
	select {}
	//select {}: 这是一个空的无限循环，它会阻塞程序的执行，直到收到退出信号或程序被终止。
	//在这里，它的作用是使程序保持运行状态，
	//以便 cron 调度器可以持续执行定时任务。
	//因为 Cron.Start() 是一个非阻塞的操作，所以需要一个无限循环来保持程序的运行，否则程序执行完后就会退出，导致定时任务无法执行。
}

func Func1() {
	fmt.Println("func1", time.Now())
	//
	//return func() {
	//
	//}
}

func inner(name string) func() {
	return func() {
		fmt.Printf("%s %s \n", name, time.Now())
	}
}

type DummyJob struct {
	Name string `json:"name"`
}

func (d DummyJob) Run() {
	fmt.Println(d.Name)
}
