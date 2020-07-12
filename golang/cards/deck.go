package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "", "Ten", "9", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of"+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}