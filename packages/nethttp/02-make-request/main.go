package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	makeRequest()
}

func makeRequest() {
	resp, err := http.Get("https://jamfpro.parsons.com/healthCheck.html")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sc := resp.StatusCode
	fmt.Println(sc)

	log.Println(string(body))
}
