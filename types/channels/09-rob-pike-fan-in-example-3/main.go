package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// print number of go routines at time of start: 1
	// only the func main go routine
	fmt.Println(runtime.NumGoroutine())
	c := fanIn(boring("Joe"), boring("Ann"), boring("Peter"), boring("Nicole"))
	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
	// print number of go routines at timie of stop: 9
	// 1 for the main func (1 go routine)
	// 1 for each fanIn go func (4 passed inputs to fanIn)
	// 1 for each boring go func (4 passed boring funcs)
	fmt.Println(runtime.NumGoroutine())
}
func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func fanIn(input1, input2, input3, input4 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	go func() {
		for {
			c <- <-input3
		}
	}()
	go func() {
		for {
			c <- <-input4
		}
	}()
	return c
}
