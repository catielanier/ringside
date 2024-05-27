package main

import "fmt"

func main() {
	deck := deck{
		Superstar: superstar{
			Name:               "Mankind",
			IsRevolution:       false,
			HandSize:           2,
			SuperStarValue:     4,
			ArsenalLimit:       60,
			BacklashLimit:      20,
			BackstageLimit:     99,
			DrawSize:           2,
			Alignment:          Any,
			SecondaryAlignment: SecondaryAny,
			NonUniqueLimit:     3,
			UniqueLimit:        1,
			EugeneMode:         false,
			CardText:           "You must always draw 2 cards, if possible, during your draw segment. All damage from opponent is at -1D",
		},
		BackstageArea: backstageArea{},
		BacklashDeck:  backlashDeck{},
		Arsenal:       newArsenal(),
		DeckType:      Virtual,
	}
	deck.Arsenal.shuffle()
	hand, remainingArsenal := deal(deck.Arsenal, deck.Superstar.HandSize)

	fmt.Println("Your superstar is: " + deck.Superstar.Name)
	fmt.Println("Your hand is:")
	for _, card := range hand {
		fmt.Println(card)
	}
	fmt.Println("Your remaining arsenal is:")
	for _, card := range remainingArsenal {
		fmt.Println(card)
	}
}
