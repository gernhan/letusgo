package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	return err
}

func (d deck) newDeckFromFile(fileName string) (deck, error) {
	inByte, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	s := strings.Split(string(inByte), ",")
	return s, nil
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
