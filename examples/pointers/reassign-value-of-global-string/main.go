package main

import (
	"fmt"
	"net/http"
)

var name = "peter"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, name)
}

func main() {
	fmt.Println(name)
	name = "nicole"
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("0.0.0.0:6060", nil)
	if err != nil {
		fmt.Println(err)
	}
}
