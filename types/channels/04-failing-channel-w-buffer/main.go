// Because the channel only has 1 buffer space and the int hasn't been pulled off, it is full.

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 8
	c <- 9

	fmt.Println(<-c)
}
