package main

import "fmt"

func main() {
	cards := deck{"A card", newCard()}
	cards = append(cards, "another card")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "a Card"
}
