package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(s)
	fmt.Printf("%p\n", s)
	s = s[:1]
	fmt.Println(s)
	fmt.Printf("%p\n", s)
	s = s[:4]
	fmt.Println(s)
	fmt.Printf("%p\n", s)
}
