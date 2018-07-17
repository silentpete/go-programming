/*

Welcome to COS Gophers! (#3)

How many people develop?

How many people have seen/worked with Go code?
- Where to start - http://blog.pg-h.io/20180610

Anyone run a website? (blog or something personal)

*/

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>COS Gophers On The Web!!!</h1>")
	fmt.Println(r.RequestURI)
}

func main() {
	fmt.Println("COS Gophers")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("0.0.0.0:6060", nil)
	if err != nil {
		fmt.Println(err)
	}
}
