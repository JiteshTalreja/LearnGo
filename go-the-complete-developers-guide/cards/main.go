package main

var cards deck

func main() {

	cards = newDeck()

	hand, remainingDeck := deal(cards, 3)

	hand.printCards()
	remainingDeck.printCards()
}
