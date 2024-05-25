package main

import (
	"fmt"
	"math/rand"
)

type alignment int
type secondaryAlignment int

const (
	Heel alignment = iota
	Face
	Neutral
	Any
)

const (
	Cheater secondaryAlignment = iota
	FanFavorite
	SecondaryNeutral
	SecondaryAny
)

type superstar = struct {
	Name               string
	IsRevolution       bool
	HandSize           int
	SuperStarValue     int
	ArsenalLimit       int
	BacklashLimit      int
	BackstageLimit     int
	DrawSize           int
	Alignment          alignment
	SecondaryAlignment secondaryAlignment
	EugeneMode         bool
	NonUniqueLimit     int
	UniqueLimit        int
	CardText           string
}

type arsenalCardType int

const (
	Maneuver arsenalCardType = iota
	Action
	Antic
	Reversal
)

type brand int

const (
	Raw brand = iota
	Smackdown
	NoBrand
)

type arsenalCard struct {
	Title              string
	IsRevolution       bool
	IsThrowback        bool
	IsSurvivorSeries   bool
	IsForeignObject    bool
	IsChain            bool
	IsHeat             bool
	IsRMS              bool
	IsUnique           bool
	IsActive           bool
	Brand              brand
	Alignment          alignment
	SecondaryAlignment secondaryAlignment
	SuperstarSpecific  string
	Volley             int
	StunValue          int
	CardType           []arsenalCardType
	ManeuverCardType   []maneuverCardType
	ManeuverSubType    maneuverSubType
}

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
	Superstar     superstar
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
