package main 

import "fmt"

type game struct {
	name  string
	price float64
}

func (b *game) print() {
	fmt.Printf("Game name: %s, price: %.3f \n", b.name, b.price)
}
