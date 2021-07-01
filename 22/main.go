package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := make([]int, 0, 20)
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < cap(x); i++ {
		x = append(x, rand.Intn(15))
	}

	fmt.Println(x)
	quickSort(x)

	fmt.Println(x)
}

func qsortPass(arr []int, done chan int) []int {
	if len(arr) < 2 {
		done <- len(arr)
		return arr
	}
	pivot := arr[0]
	i, j := 1, len(arr)-1
	for i != j {
		for arr[i] < pivot && i != j {
			i++
		}
		for arr[j] >= pivot && i != j {
			j--
		}
		if arr[i] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if arr[j] >= pivot {
		j--
	}
	arr[0], arr[j] = arr[j], arr[0]
	done <- 1

	go qsortPass(arr[:j], done)
	go qsortPass(arr[j+1:], done)
	return arr
}

func quickSort(arr []int) {
	done := make(chan int)

	go qsortPass(arr, done)

	rslt := len(arr)
	for rslt > 0 {
		rslt -= <-done
	}
	return
}
