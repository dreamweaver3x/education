package main

import (
	"sync/atomic"
)

func main() {
	var x int32
	var y int32

	x = 5
	y = 10

	y = atomic.SwapInt32(&x, y)
	println(x, y)

	x, y = y, x
	println(x, y)
}
