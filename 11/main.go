package main

import "fmt"

func main() {
	x := [10]int{9, 8, 1, 4, 1, 2, 0, 11, 19, 32}
	y := [8]int{32, 1, 8, 80, 8, 8, 1, 19}
	mapa := make(map[int]uint8)

	for _, v := range x {
		mapa[v] = 0
	}

	for _, v := range y {
		num, ok := mapa[v]
		if ok {
			if num == 0 {
				fmt.Println(v)
				mapa[v]++
			}
		}
	}
}
