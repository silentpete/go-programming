package main

import (
	"fmt"
	"strings"
)

func main() {
	host := "server.domain.tld:1234"
	fmt.Println(strings.TrimRight(host, ":1234567890"))
}
