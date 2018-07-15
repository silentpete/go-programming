// Reference: http://www.golang-book.com/books/intro/13
// To open a file in Go use the Open function from the os package.
// Here is an example of how to read the contents of a file and display them on the terminal:

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
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

	str := string(bs)
	// print out the file contents
	fmt.Println(str)
}
