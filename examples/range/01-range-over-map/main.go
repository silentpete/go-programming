package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	byt := []byte(`{"name":"john","num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	for k, v := range dat {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println("", i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
