package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go workerOne(ch1, ch2)
	go workerTwo(ch2)

	for i := 0; i < 1000; i++ {
		ch1 <- rand.Intn(100)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("done")
}

func workerOne(ch1, ch2 chan int) {
	var x int
	for {
		x = <-ch1
		if x%2 == 0 {
			ch2 <- x
		}
	}
}

func workerTwo(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
