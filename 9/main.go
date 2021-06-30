package main

import (
	"fmt"
	"time"
)

func main() {
	x := [100]int{}
	for i := 0; i < 100; i++ {
		x[i] = i
	}
	chOne := make(chan int)
	chTwo := make(chan int)
	go workerOne(chOne, chTwo)
	go workerTwo(chTwo)
	for i := 0; i < 100; i++ {
		fmt.Printf("x[%d] = %d\n", i, x[i])
		chOne <- x[i]
		time.Sleep(time.Millisecond * 200)
	}
}

func workerOne(chOne, chTwo chan int) {
	for {
		chTwo <- <-chOne * 2
	}
}

func workerTwo(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
