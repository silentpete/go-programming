package main

import (
	"fmt"
	"log"
)

func main() {
	var userInput string
	fmt.Println("please enter something and press Enter:")
	i, err := fmt.Scanln(&userInput)
	if i <= 0 {
		log.Fatal("nothing entered")
	}
	if err != nil {
		log.Fatal("ERROR: Scanln issue:", err)
	}
	log.Println("received:", userInput)
}
