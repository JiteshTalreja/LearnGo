package main

import "fmt"

/* Create a new dataype called deck
deck will inherit propeties or will be based on Slice
*/

type deck []string

func (d deck) printCards() {

	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}

	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
