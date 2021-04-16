package main 

import "fmt"

type printer interface {
	print()
}

// abstract layer for book, game, and movie type
type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("nothing to be print then")
		return 
	}
	for _, it := range l {
		fmt.Printf("type : %-10T -->", it)
		it.print()
	}
}