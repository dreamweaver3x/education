package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	var sec int

	fmt.Print("insert time to work(seconds): ")
	fmt.Scan(&sec)

	workTime := time.After(time.Second * time.Duration(sec))

	go worker(ch)
LOOP:
	for {
		select {
		case <-workTime:
			break LOOP
		default:
			ch <- rand.Intn(100)
		}
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println("time's up")
}

func worker(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
