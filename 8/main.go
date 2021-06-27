package main

import "fmt"

func main() {
	var x, y int64
	var bit, change int
	fmt.Print("insert value: ")
	fmt.Scan(&x)
	fmt.Printf("this is how %d looks in bits: %b\n", x, x)

	fmt.Print("insert which bit u want to change on on what :")
	fmt.Scan(&bit, &change)
	y = 1 << (bit - 1)

	if change == 1 {
		fmt.Printf("y = %d, in bits: %b\n", y, y)
		x = x | y
	} else if change == 0 {
		fmt.Printf("y = %d, in bits: %b\n", y, y)
		x = x &^ y
	} else {
		fmt.Println("new bit should be 0 or 1")
		return
	}

	fmt.Printf("new value is %d, looks in bits: %b\n", x, x)
}
