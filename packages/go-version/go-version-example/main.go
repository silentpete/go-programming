package main

import (
	"fmt"

	"github.com/mcuadros/go-version"
)

func main() {
	fmt.Println("testing mcuadros/go-version")

	v1 := "1.0.0"
	v2 := "1.0.1"

	v3 := "2.7.1.27107788"
	v4 := "2.4.1.27107788"

	v5 := "1.0.39-jira8"
	v6 := "1.0.34-jira8"

	v7 := "6.5.0-p5"
	v8 := "6.4.1-p5"

	fmt.Println("is", v1, "less than", v2, "? ", version.Compare(v1, v2, "<"))
	fmt.Println("is", v3, "less than", v4, "? ", version.Compare(v3, v4, "<"))
	fmt.Println("is", v5, "less than", v6, "? ", version.Compare(v5, v6, "<"))
	fmt.Println("is", v7, "less than", v8, "? ", version.Compare(v7, v8, "<"))
}
