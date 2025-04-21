package main

import (
	"sync"
	"time"
)

func task1() {
	time.Sleep(time.Second * 5)
	println("task1执行完毕")
}
func task2() {
	time.Sleep(time.Second * 3)
	println("task2执行完毕")
}
func asynctask1(event *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	println("task1执行完毕")
	event.Done()
}
func asynctask2(event *sync.WaitGroup) {
	time.Sleep(time.Second * 3)
	println("task2执行完毕")
	event.Done()
}
func main() {
	// 串行执行
	// 耗时：8秒
	start := time.Now().UnixMilli()
	task1()
	task2()
	end := time.Now().UnixMilli()
	println("串行总耗时：", end-start)

	// 并发执行
	// 耗时：5秒
	var event = sync.WaitGroup{}
	start = time.Now().UnixMilli()
	event.Add(2)
	go asynctask1(&event)
	go asynctask2(&event)
	event.Wait()
	end = time.Now().UnixMilli()
	println("串行总耗时：", end-start)
}
