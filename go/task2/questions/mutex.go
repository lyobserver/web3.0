package questions

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func MutexRun() {
	var count int64 = 0
	var mutex sync.Mutex
	var event sync.WaitGroup
	event.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer event.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	event.Wait()
	fmt.Println("使用mutex累计count的值为：", count)
}

func AtomicRun() {
	var count int64 = 0
	var event sync.WaitGroup
	event.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer event.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	event.Wait()
	fmt.Println("使用atomic累计count的值为：", count)
}
