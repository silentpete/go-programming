package main

import "fmt"

type shoe struct {
	ID    int     `json:"id"`
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Color string  `json:"color"`
	Size  int     `json:"size"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func main() {
	var someShoe shoe
	var shoes []shoe

	someShoe = shoe{
		ID:    0,
		Brand: "nike-men",
		Model: "air force one",
		Color: "white",
		Size:  10,
		Price: 99.99,
		Stock: 0,
	}
	shoes = append(shoes, someShoe)

	fmt.Println(shoes)
}
