package questions

import (
	"fmt"
	"sync"
	"time"
)

func OddEvenRun() {
	var event sync.WaitGroup
	event.Add(2)
	go odd(&event)
	go even(&event)
	event.Wait()
}

func odd(e *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("奇数：", i)
		}
		time.Sleep(time.Second)
	}
	e.Done()
}

func even(e *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("偶数：", i)
		}
		time.Sleep(time.Second)
	}
	e.Done()
}

type result struct {
	index    int
	duration time.Duration
}

func SchedulerRun() {
	tasks := []func(){
		func() { time.Sleep(1 * time.Second) },
		func() { time.Sleep(2 * time.Second) },
		func() { time.Sleep(3 * time.Second) },
	}
	results := scheduler(&tasks)
	for _, result := range results {
		fmt.Printf("任务%d执行时间：%v\n", result.index, result.duration)
	}
}

func scheduler(tasks *[]func()) []result {
	var results []result
	var wg sync.WaitGroup
	wg.Add(len(*tasks))
	for i, task := range *tasks {
		go func(i int, task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			results = append(results, result{index: i, duration: duration})
		}(i, task)
	}
	wg.Wait()
	return results
}
