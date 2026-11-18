package main

import "fmt"

var cards []string

func main() {

	cards = append(cards, "ace of spades", newCard(), "queen of hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {

	return "five of spades"
}
