package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Peter"
	if !strings.Contains("", name) {
		fmt.Println("Hello", name)
	}
}
