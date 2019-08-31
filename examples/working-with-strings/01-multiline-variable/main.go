package main

import "fmt"

func main() {
	multilineStringOne := "this is line one " +
		"this is line two"
	multilineStringTwo := "this is line one\n" +
		"this is line two"
	fmt.Println(multilineStringOne)
	fmt.Println(multilineStringTwo)
}
