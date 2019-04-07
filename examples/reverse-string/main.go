// https://youtu.be/XCsL89YtqCs
package main

import "fmt"
// For writing example code, this is the short way to the directory with a sub package.
import "./mystring"

func main() {
	fmt.Println("Hello World")
	fmt.Println(mystring.Reverse("Hello World"))
}
