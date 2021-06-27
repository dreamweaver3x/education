package main

import "fmt"

func main() {
	var y int

	fmt.Print("insert slice len: ")
	fmt.Scan(&y)

	x := make([]int, y)
	for i := 0; i < y; i++ {
		x[i] = i + 1
	}

	fmt.Println(x)
	fmt.Print("insert index to delete: ")
	fmt.Scan(&y)

	x[y-1] = x[len(x)-1]
	x = x[:len(x)-1]

	fmt.Println(x)
}
