package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// have to close the channel or range will stay open forever
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
