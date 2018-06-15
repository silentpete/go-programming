package main

import "fmt"

func main() {
	fmt.Print("Enter a number that will be multiplied by two: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
