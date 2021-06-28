package main

import (
	"fmt"
	"sync"
)

type mapka struct {
	sync.Mutex
	x map[int]int
}

func main() {
	test := mapka{
		x: make(map[int]int),
	}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			test.Lock()
			test.x[i%11]++
			test.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	for k, v := range test.x {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}
}
