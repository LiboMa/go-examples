package main 

import "fmt"

// game type in the store definition 
type game struct {
	name  string
	price float64
}

func (b *game) print() {
	fmt.Printf("Game name: %s, price: %.3f \n", b.name, b.price)
}
