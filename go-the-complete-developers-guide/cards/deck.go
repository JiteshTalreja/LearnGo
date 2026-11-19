package main

import "fmt"

/* Create a new dataype called deck
deck will inherit propeties or will be based on Slice
*/

type deck []string

func (d deck) printCards() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
