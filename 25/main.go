package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type counter struct {
	sync.Mutex
	x int
}

func main() {
	count := counter{}

	for i := 0; i < 100; i++ {
		go func() {
			count.Lock()
			count.x++
			fmt.Println("counter = ", count.x)
			count.Unlock()
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("final counter = ", count.x)

	var x int32
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 100; c++ {
				atomic.AddInt32(&x, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("x = ", x)

}
