package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := deck{"Spades", "Diamonds", "Hearts"}
	cardValues := deck{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return []string{}
}

func (d deck) print() {
	for i, d := range d {
		fmt.Println(i, d)
	}
}
