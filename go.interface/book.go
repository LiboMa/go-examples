package main 

import (
	"fmt"
)

// book type definition 
type book struct {
	name  string
	price float64
	kind  string
}

func (b *book) print() {
	fmt.Printf("the book name: %s, price: %.3f \n", b.name, b.price)
}
