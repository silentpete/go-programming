// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println(cap(a))
}
