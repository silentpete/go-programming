package main

import (
	"fmt"
)

func main() {
	// bidirectional channel
	c := make(chan int)
	// send channel
	cs := make(chan<- int)
	// receive channel
	cr := make(<-chan int)

	go func() {
		c <- 8
	}()

	// bidirectional can be converted to a send or receive,
	// but a send or receive channel CANNOT be converted to a bidirectional channel
	fmt.Println(<-c)
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cr\t%T\n", cs)
}
