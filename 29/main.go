package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(20000000000)
	y := big.NewInt(10000000000)
	z := new(big.Int)
	fmt.Println(x)
	fmt.Println(y)
	z.Mul(x, y)
	fmt.Println("multiplication = ", z)
	z.Add(x, y)
	fmt.Println("add = ", z)
	z.Sub(x, y)
	fmt.Println("sub =", z)
	z.Div(x, y)
	fmt.Println("div = ", z)
}
