// Reference: http://www.golang-book.com/books/intro/13
// To open a file in Go use the Open function from the os package.
// Here is an example of how to read the contents of a file and display them on the terminal:

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("many")
	if err != nil {
		// handle the error here
		log.Fatal(err)
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	pattern := "Client Name"
	patternMatch := regexp.MustCompile(pattern)

	linesAry := strings.Split(string(bs), "\n")
	lineAndCountAry := make(map[string]int)
	for _, line := range linesAry {
		if patternMatch.MatchString(line) {
			if lineAndCountAry[strings.TrimSpace(line)] == lineAndCountAry[strings.TrimSpace(line)] {
				lineAndCountAry[strings.TrimSpace(line)]++
			}
			if lineAndCountAry[strings.TrimSpace(line)] != lineAndCountAry[strings.TrimSpace(line)] {
				lineAndCountAry[strings.TrimSpace(line)] = 1
			}
		}
	}

	cnt := 0
	for _, v := range lineAndCountAry {
		cnt = cnt + v
	}
	fmt.Println("total matches:", cnt)

	fmt.Println("unique count:", len(lineAndCountAry))

	fmt.Println(lineAndCountAry)

}
