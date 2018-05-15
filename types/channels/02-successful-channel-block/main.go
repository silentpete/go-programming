// This will NOT run
// A channel "blocks"
package main

import (
	"fmt"
)

func main() {
	// create channel of type int
	c := make(chan int)
	// put something on the channel
	go func() {
		c <- 8
	}()
	// print the channel
	fmt.Println(<-c)
}
