package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) { //нужно передавать указатель на waitGroup, иначе создается новая
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
