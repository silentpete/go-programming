package main

// Create non compliant code for SonarQube example

// Noncompliant Code Example
// if boolFunc() == true {
//     // ...
// }
// flag := x && true

// Compliant Solution
// if boolFunc() {
//     // ...
// }
// flag := x

import (
	"fmt"
)

// Noncompliant
func main() {
	x := false
	if boolFunc() == true {
		fmt.Println("true")
	}
	flag := x && true
	fmt.Println(flag)
}

// SWAP and check with SonarQube

// Compliant
// func main() {
// 	x := false
// 	if boolFunc() {
// 		fmt.Println("true")
// 	}
// 	flag := x
// 	fmt.Println(flag)
// }

func boolFunc() bool {
	return true
}
