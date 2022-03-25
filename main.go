package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 7)

	hand.print()
	remainingCards.print()
}
