package main

import "fmt"

const someConst = 12

// can initialize a variable, but can not assign a value to it
var someVar int

func main() {
	// initialization of card
	card := i2S(newCard())
	// card := "Ace of Spades"
	//var card = "Ace of Spades"
	//var card string = "Ace of Spades"

	// assign variable in the next times
	//card = "Five of Diamonds"
	cards := deck{card, "Ace of Spades"}
	cards = append(cards, "Six of Spades")

	//fmt.Println(len(card))
	cards.print()

	cards = newDeck()
	cards.print()
}

func i2S(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

func newCard() interface{} {
	return "Five of Diamonds"
}
