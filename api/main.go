package main

import (
	"fmt"
	"github.com/catielanier/ringside/cards"
	"github.com/catielanier/ringside/cards/global"
	"github.com/catielanier/ringside/configs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiURI := os.Getenv("API_URI")
	deck := deck{
		Superstar: cards.Superstar{
			Name:               "Mankind",
			IsRevolution:       false,
			HandSize:           2,
			SuperStarValue:     4,
			ArsenalLimit:       60,
			BacklashLimit:      20,
			BackstageLimit:     99,
			DrawSize:           2,
			Alignment:          global.Any,
			SecondaryAlignment: global.SecondaryAny,
			NonUniqueLimit:     3,
			UniqueLimit:        1,
			EugeneMode:         false,
			CardText:           "You must always draw 2 cards, if possible, during your draw segment. All damage from opponent is at -1D",
		},
		BackstageArea: []cards.Backstage{},
		BacklashDeck:  []cards.Backlash{},
		Arsenal:       arsenal{},
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

	router := gin.Default()

	configs.ConnectDB()

	router.Run(apiURI)
}
