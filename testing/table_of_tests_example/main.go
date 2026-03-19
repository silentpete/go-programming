package main

import (
	"fmt"
)

func main() {
	fmt.Println("the sum:", sum(1, 2, 3))

	username := "admin"
	password := "password123"

	// Simulate using credentials
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

}

func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum = sum + v
	}
	return sum
}
