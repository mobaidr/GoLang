package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	cards.shuffle()

	fmt.Println("===========================================")

	cards.print()
	//cards.saveToFile("Cards.txt")

}
