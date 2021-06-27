package main

import (
	"fmt"
	"sync"
)

func main() {
	x := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, v := range x {
		wg.Add(1)
		go func(v int) {
			fmt.Printf("%d squared equals to %d\n", v, v*v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
