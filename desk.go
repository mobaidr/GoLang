package main

import "fmt"

//Create a new type o Deck which is slice of string
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
