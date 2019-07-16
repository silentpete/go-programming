package main

import (
	"fmt"
	"time"
)

var (
	startTime    time.Time
	timeNow      time.Time
	waitDuration = time.Duration(1 * time.Minute)
)

func main() {
	startTime = time.Now()

	loop := 18
	for index := 0; index < loop; index++ {
		timeNow = time.Now()
		fmt.Println("working...")
		if timeNow.Sub(startTime) >= waitDuration {
			fmt.Println("it's time for a break")
			startTime = time.Now()
		}
		time.Sleep(time.Duration(10 * time.Second))
	}
	fmt.Println("enough working, I'm done.")
}
