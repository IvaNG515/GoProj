package main

func main() {
	cards := newDeck()

	hand, leftCards := deal(cards, 10)
	hand.print()
	leftCards.print()
}
