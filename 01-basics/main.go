package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("cards")

	newCards := newDeckFromFile("cards")
	fmt.Println(newCards.toString())

	newCards.shuffle()

	newCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
