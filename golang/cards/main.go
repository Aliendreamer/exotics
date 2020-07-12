package main

import "fmt"

func main() {
	cards := newDeck()
	deckOne, deckTwo := deal(cards, 5)
	deckOne.print()
	fmt.Println("**********")
	deckTwo.print()
}
