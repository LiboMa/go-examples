package main 

import (
	"fmt"
)

type book struct {
	name  string
	price float64
	kind  string
}

func (b *book) print() {
	fmt.Printf("the book name: %s, price: %.3f \n", b.name, b.price)
}
