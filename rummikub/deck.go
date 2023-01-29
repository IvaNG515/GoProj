package main

import (
	"fmt"
	"math/rand"
	"time"
)

type full_deck []string

func newDeck() full_deck {
	// deck will automatically be randomized on newDeck()
	full_deck := full_deck{}
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
	// randomization
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range full_deck {
		newPos := r.Intn(len(full_deck) - 1)
		full_deck[i], full_deck[newPos] = full_deck[newPos], full_deck[i]
	}

	return full_deck
}

func (d full_deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
