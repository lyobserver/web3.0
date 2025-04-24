package questions

import (
	"fmt"
	"sync"
	"time"
)

func ChannelRun() {
	var ch = make(chan int)
	var e sync.WaitGroup
	e.Add(2)
	go writeNumber(&ch, &e)
	go readNumber(&ch, &e)
	e.Wait()
}

func writeNumber(ch *chan int, e *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		*ch <- i
		time.Sleep(500)
	}
	close(*ch)
	e.Done()
}

func readNumber(ch *chan int, e *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		num := <-*ch
		fmt.Println(num)
	}
	e.Done()
}

func CacheChannelRun() {
	var ch = make(chan int, 100)
	var e sync.WaitGroup
	go cacheWriteNumber(&ch)
	e.Add(1)
	go cacheReadNumber(&ch, &e)
	e.Wait()
}

func cacheWriteNumber(ch *chan int) {
	for i := 0; i < 100; i++ {
		*ch <- i
	}
}

func cacheReadNumber(ch *chan int, e *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		num := <-*ch
		fmt.Println(num)
	}
	e.Done()
}
