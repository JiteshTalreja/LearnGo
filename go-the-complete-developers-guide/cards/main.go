package main

var cards deck

func main() {

	cards = append(cards, "ace of spades", newCard(), "queen of hearts")

	cards.printCards()
}

func newCard() string {

	return "five of spades"
}
