package random

import (
	"fmt"
)

// TestPrintHelloWorld is a simple test for testing the return of a function.
func ExampleReturnHelloWorld() {
	fmt.Println(ReturnHelloWorld())
	// Output:
	// hello world
}

func ExampleReturnPassedString() {
	fmt.Println(ReturnPassedString("peter"))
	// Output:
	// peter
}
