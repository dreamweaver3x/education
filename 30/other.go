package main

import "fmt"

func main() {
	var y int
	var newSlice []int

	fmt.Print("insert slice len: ")
	fmt.Scan(&y)

	x := make([]int, y)
	for i := 0; i < y; i++ {
		x[i] = i + 1
	}

	fmt.Println(x)
	fmt.Print("insert index to delete: ")
	fmt.Scan(&y)

	firstHalf := x[:y-1]
	if y < len(x) {
		secondHalf := x[y:]
		newSlice = append(firstHalf, secondHalf...)
	} else {
		newSlice = firstHalf
	}

	fmt.Println(newSlice)
}
