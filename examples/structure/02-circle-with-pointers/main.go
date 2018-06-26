package main

import (
	"fmt"
	"math"
)

type circle struct {
	x float64
	y float64
	r float64
}

func circleArea(c *circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := circle{0, 0, 5}
	fmt.Println(circleArea(&c))
}
