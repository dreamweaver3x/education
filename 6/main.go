package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	ch := make(chan struct{})
	ctxOne, _ := context.WithTimeout(context.TODO(), time.Second*5)
	ctxTwo, cancel := context.WithCancel(context.TODO())

	for i := 0; i < 2; i++ {
		go goMutex(&mu, i)
		go goWaitGroup(&wg, i)
		go goChan(ch, i)
	}
	go goTimeOut(ctxOne)
	go goCancel(ctxTwo)
	go goSleep()

	ch <- struct{}{}
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 5)

}

func goMutex(mu *sync.Mutex, i int) {
	fmt.Printf("go mutex %d started\n", i)
	mu.Lock()
	fmt.Printf("go mutex %d ended\n", i)
}

func goWaitGroup(wg *sync.WaitGroup, i int) {
	fmt.Printf("go waitGroup %d started\n", i)
	wg.Wait()
	wg.Add(1)
	fmt.Printf("go waitGroup %d ended\n", i)
}

func goChan(ch chan struct{}, i int) {
	fmt.Printf("go chan %d started\n", i)
	<-ch
	fmt.Printf("go chan %d ended\n", i)
}

func goTimeOut(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("go timeOut ended\n")
		default:
			fmt.Printf("go timeOut is working\n")
		}
		time.Sleep(time.Second)
	}
}

func goCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("go cancel ended\n")
		default:
			fmt.Printf("go cancel is working\n")
		}
		time.Sleep(time.Second)
	}
}

func goSleep() {
	fmt.Printf("go sleep started\n")
	time.Sleep(time.Second * 4)
	fmt.Printf("go sleep ended\n")
}
