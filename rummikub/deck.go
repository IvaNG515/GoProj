package main

import "fmt"

type deck []string

func newDeck() deck {
	full_deck := deck{}
	cardNum := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}
	cardColor := []string{"R", "Y", "B", "BL"}
	cnt := 0
	for cnt != 2 {
		for _, color := range cardColor {
			for _, num := range cardNum {
				full_deck = append(full_deck, num+color)
			}
		}
		cnt += 1
	}
	full_deck = append(full_deck, "Joker1")
	full_deck = append(full_deck, "Joker2")

	return full_deck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
