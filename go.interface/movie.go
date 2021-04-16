package main 

import "fmt"

type movie struct {
	name  string
	price float64
}

// movie type in the store definition 

func (b *movie) print() {
	fmt.Printf("the movie name: %s, price: %.3f \n", b.name, b.price)
}
