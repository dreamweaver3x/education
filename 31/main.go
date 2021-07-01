package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func main() {
	pointOne := point{}
	pointTwo := point{}

	fmt.Print("insert x and y coordinates for point one: ")
	fmt.Scan(&pointOne.x, &pointOne.y)
	fmt.Println()
	fmt.Print("insert x and y coordinates for point two: ")
	fmt.Scan(&pointTwo.x, &pointTwo.y)

	fmt.Println(pointOne, pointTwo)

	dist := pointOne.distance(pointTwo)
	fmt.Printf("distance = %v", dist)

}

func (p point) distance(p2 point) float64 {
	return math.Sqrt(float64((p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y)))
}
