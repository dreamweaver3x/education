package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100       //слайс внутри функции указывает на ту же область памяти, куда указывает слайс a
	v = append(v, b) //теперь v указывает на другую область памяти, не связанной со слайсом a
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println("len a = ", len(a), "cap a = ", cap(a))
	fmt.Println(a)
}
