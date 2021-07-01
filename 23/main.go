package main

import "fmt"

func main() {
	x := make([]int, 50)

	for i, _ := range x {
		x[i] = i * 2
	}

	fmt.Println(x)
	find := 98
	res, count := binarySearch(x, find)
	fmt.Printf("found number %d on %d slot in %d iterations\n", find, res, count)
}

func binarySearch(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1
	case a[mid] > search:
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearch(a[mid+1:], search)
		if result >= 0 {
			result += mid + 1
		}
	default:
		result = mid
	}
	searchCount++
	return
}
