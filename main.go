package main

func main() {
	//cards := newDeck()

	//cards.saveToFile("Cards.txt")

	newDeck := newDeckFrom("Cards.txt")

	newDeck.print()
}
