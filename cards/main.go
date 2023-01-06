package main

func main() {
	cards := deck{"A card", newCard()}
	cards = append(cards, "another card")

	cards.print()
}

func newCard() string {
	return "a Card"
}
