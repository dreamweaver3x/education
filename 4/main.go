package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	var workers int
	var data string
	ctx, cancel := context.WithCancel(context.TODO())

	fmt.Print("insert num of workers: ")
	fmt.Scan(&workers)

	for i := 0; i < workers; i++ {
		go work(i, ch, ctx)
	}

	fmt.Print("insert data, workers will take it: ")
	for {
		fmt.Scanf("%s\n", &data)
		if data == "cancel" {
			cancel()
			break
		}
		ch <- data
		time.Sleep(time.Millisecond * 200)
	}

}

func work(i int, ch chan string, ctx context.Context) {
	var data string
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case data = <-ch:
			fmt.Printf("worker num %d took data: %q\n", i+1, data)
		default:
		}
	}
}
