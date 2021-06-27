package main

import "fmt"

func main() {
	n := 0
	if true { // внутри if инициализируется локальная переменная, ее изменения не связаны с внешней переменной n
		n := 1
		n++
	}
	fmt.Println(n)
}
