package main

import (
	"fmt"
)

func main() {
	var array1 = []string{"abc", "defg", "hijkl"}
	fmt.Println("array length:", len(array1))
	x := 0
	for index, value := range array1 {
		fmt.Println("printing value", index, "length:", len(value))
		x = len(value) + x
	}
	fmt.Println("add up all letters:", x)
}
