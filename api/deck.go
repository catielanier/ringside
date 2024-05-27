package main

import (
	"fmt"
	"github.com/catielanier/ringside/cards"
	"math/rand"
)

type backstageArea []string

type backlashDeck []string

type arsenal []string

type deckType int

const (
	Classic deckType = iota
	Virtual
	Revolution
)

type deck struct {
	Superstar     cards.Superstar
	BackstageArea backstageArea
	BacklashDeck  backlashDeck
	Arsenal       arsenal
	DeckType      deckType
}

func newArsenal() arsenal {
	return arsenal{"Escape Move", "Step Aside", "Right Cross Punch", "Volley This!", "Reach for the Ropes", "There Are Two Things You Can Do: Nothing and Like It"}
}

func (a arsenal) print() {
	for i, card := range a {
		fmt.Println(i, card)
	}
}

func (a arsenal) shuffle() {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func deal(a arsenal, handSize int) (arsenal, arsenal) {
	return a[:handSize], a[handSize:]
}
