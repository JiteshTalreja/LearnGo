package main

import "fmt"

var cards deck

func main() {

	cards = newDeck()

	hand, remainingDeck := deal(cards, 3)

	hand.printCards()
	remainingDeck.printCards()

	fmt.Println([]byte(hand.toString()))
}
