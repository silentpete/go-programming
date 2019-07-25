package main

import (
	"fmt"
	"strings"
)

func main() {
	name := " Peter Nicole "
	name = strings.Replace(name, " ", "", -1) // remove all whitespace
	fmt.Printf("'%s'", name)
}
