// from go tour
package main

import "fmt"

func main() {
	sum := 1
	for i := 0; i < 9; i++ {
		sum += sum
	}
	fmt.Println(sum)
}
