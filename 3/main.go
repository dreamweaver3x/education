package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	x := [5]int{2, 4, 6, 8, 10}
	var sum int32
	wg := sync.WaitGroup{}

	for _, v := range x {
		wg.Add(1)
		go func(v int) {
			sum = atomic.AddInt32(&sum, int32(v*v))
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println("sum of squares = ", sum)
}
