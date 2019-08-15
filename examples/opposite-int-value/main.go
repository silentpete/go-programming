package main

import "fmt"

var (
	t1 = 1
	t2 = 2
	t3 = -3
	t4 = 0
)

func main() {
	fmt.Println(opposite(t1))
	fmt.Println(opposite(t2))
	fmt.Println(opposite(t3))
	fmt.Println(opposite(t4))
}

func opposite(value int) int {
	return -value
}
