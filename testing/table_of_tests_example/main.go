package main

import (
	"fmt"
)

func main() {
	fmt.Println("the sum:", sum(1, 2, 3))
}

func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum = sum + v
	}
	return sum
}
