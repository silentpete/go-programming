package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("boring!") // Function returning a channel.
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		// why doesn't this work? hum
		c <- fmt.Sprintf("%s", msg)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

	}()
	return c // Return the channel to the caller.
}
