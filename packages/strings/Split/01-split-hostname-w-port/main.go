package main

import (
	"fmt"
	"strings"
)

func main() {
	host := "server.domain.tld:1234"
	fmt.Println("before:", host)
	fmt.Println("after:", strings.Split(host, ":")[0])
}
