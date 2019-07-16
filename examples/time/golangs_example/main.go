package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	t := time.Now()

	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	fmt.Printf("%T\n", elapsed.Seconds)
}
