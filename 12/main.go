package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p) //в go все передается по значению, мы поменяли адрес у копии, соотв p указывает на ту же область памяти, где лежит 1
	fmt.Println(*p)
}
