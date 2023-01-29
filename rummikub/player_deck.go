package main

import "fmt"

type player_deck []string

func (p player_deck) print() {
	for i, pc := range p {
		fmt.Println(i, pc)
	}
}
