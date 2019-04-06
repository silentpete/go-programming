// https://golang.org/pkg/flag/
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help = flag.Bool("help", false, "Will display this helpful dialog output.")
	name = flag.String("name", "no-name", "Use the name option to set the name.")
)

func main() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println("in the program")

	fmt.Println("Hello " + *name)
}
