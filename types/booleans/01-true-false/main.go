package main

import "fmt"

func main() {
	a := true
	b := false

	if a == b {
		fmt.Println("a equals b")
	}

	if a != b {
		fmt.Println("a is not equal to b")
	}
}
