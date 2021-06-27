package main

import (
	"fmt"
	"sync"
)

const len int = 10

func main() {
	x := [len]int{}
	wg := sync.WaitGroup{}

	for i := 0; i < len; i++ {
		x[i] = i
	}

	for _, v := range x {
		wg.Add(1)
		go func(v int) {
			fmt.Println(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
