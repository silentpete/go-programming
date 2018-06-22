package main

import "fmt"

// A function like this together with the non-local variables it references is
// known as a closure. In this case increment and the variable x form the closure.

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

// REFERENCE: http://www.golang-book.com/books/intro/7#section4
