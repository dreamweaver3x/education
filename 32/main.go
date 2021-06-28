package main

import (
	"fmt"
	"time"
)

func main() {
	sleep := 5
	mySleep(sleep)
	fmt.Println("awake")
}

func mySleep(n int) {
	ch := time.After(time.Second * time.Duration(n))
LOOP:
	for {
		select {
		case <-ch:
			break LOOP
		default:

		}
	}
}
