package cards

import "fmt"

const someConst = 12

// can initialize a variable, but can not assign a value to it
var someVar int

func main() {
	// initialization of card
	card := i2S()
	// card := "Ace of Spades"
	//var card = "Ace of Spades"
	//var card string = "Ace of Spades"

	// assign variable in the next times
	//card = "Five of Diamonds"
	cards := []string{i2S(newCard())}

	fmt.Println(len(card))
}

func i2S(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

func newCard() interface{} {
	return "Five of Diamonds"
}
