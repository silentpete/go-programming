// Can use this as a template for adding the ability to pass ?key=value and parse it to use later
//
// Example
// curl 'http://127.0.0.1:6060?email=first.last@domain.com&username=first.last&firstname=First&lastname=Last&name=First%20Last'
//
// References
// https://golang.org/pkg/net/http/
// https://en.wikipedia.org/wiki/Query_string
// https://www.urldecoder.io/golang/
package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {

	requestedURL, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		fmt.Println("experienced a problem parsing the request uri", err)
	}

	queryStrings := requestedURL.Query()

	fmt.Println("Query Strings: ")
	for key, value := range queryStrings {
		fmt.Printf("  %v = %v\n", key, value)
		if key == "email" {
			fmt.Println("we found a email: ", value)
			fmt.Fprintf(w, "we found a email: %v", value)
		}
	}

}

func main() {
	address := "0.0.0.0"
	port := "6060"
	fmt.Println("starting http server at: " + address + ":" + port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(address+":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
