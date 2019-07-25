package main

import (
	"fmt"
	"strings"
)

func main() {
	name := " Peter Nicole "
	name = strings.Trim(name, " ") // remove whitespace
	fmt.Printf("'%s'", name)
}
