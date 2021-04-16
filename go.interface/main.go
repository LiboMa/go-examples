package main

var ( docs = "show interface usage")

func main () {
	// create sample data
	oldMan := book{name: "book of old man and sea", price: 15.5}
	oldGame := game{name: "Game of star war", price: 20.34}
	oldGame2 := game{name: "Game of start war 2", price: 22.34}
	movie := movie{name: "Movie of Iorn Man II", price: 88.88}

	var l list

	// use abstract interface to print information.
	l = append(l, &oldGame, &oldGame2, &oldMan, &movie)
	l.print()

}