package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
}

func main() {
	var peter person
	fmt.Println(peter)
	peter.Firstname = "Peter"
	fmt.Println(peter)
	peter.Lastname = "Gallerani"
	fmt.Println(peter)
}
