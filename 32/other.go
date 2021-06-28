package main

import (
	"fmt"
	"time"
)

func main() {
	sleep := 5
	mySleepNew(sleep)
	fmt.Println("awake")
}

func mySleepNew(n int) {
	select {
	case <-time.After(time.Second * time.Duration(n)):
		return
	}
}
