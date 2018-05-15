package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send
	go send(c)
	// receive
	receive(c)

	fmt.Println("exiting")
}

func send(c chan<- int) {
	c <- 8
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
