package main

import (
	"fmt"
	"math/rand"
)

type deck []string

func newDeck() deck {
	return deck{"Escape Move", "Step Aside", "Right Cross Punch", "Volley This!", "Reach for the Ropes", "There Are Two Things You Can Do: Nothing and Like It"}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
