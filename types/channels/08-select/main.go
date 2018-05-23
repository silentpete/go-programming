package main

import (
	"fmt"
)

func main() {
	// make even bidirectional channel
	even := make(chan int)
	// make odd bidirectional channel
	odd := make(chan int)
	// make quit bidirectional channel
	quit := make(chan int)

	// create go routines for send function
	go send(even, odd, quit)

	// since the go routines are out there running, take things off the channels
	receive(even, odd, quit)

	fmt.Println("exiting")
}

// receive takes each bidirectional channel and assigns it to a receive direction channel
func receive(e, o, q <-chan int) {
	// infinite for loop, will run as long as there are channels out there, will return after getting a "q" channel
	for {
		select {
		case v := <-e:
			fmt.Println("even:", v)
		case v := <-o:
			fmt.Println("odd:", v)
		case v := <-q:
			fmt.Println("quit:", v)
			return
		}
	}
}

// send takes each bidirectional channel and assigns it to a send direction channel
func send(e, o, q chan<- int) {
	// how many channels, and based on modulus
	for i := 0; i < 10; i++ {
		// if the number divided by 2 has 0 remainder
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// once the loop has finished, put 0 on the q channel and the func will return
	q <- 0
}
