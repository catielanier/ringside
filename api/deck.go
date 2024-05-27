package main

import (
	"fmt"
	"github.com/catielanier/ringside/cards"
	"math/rand"
)

type arsenal []cards.ArsenalCard

type deckType int

const (
	Classic deckType = iota
	Virtual
	Revolution
)

type deck struct {
	Superstar     cards.Superstar
	BackstageArea []cards.Backstage
	BacklashDeck  []cards.Backlash
	Arsenal       arsenal
	DeckType      deckType
}

func newArsenal() arsenal {
	return arsenal{}
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
