package controllers

import (
	"context"
	"github.com/catielanier/ringside/cards"
	"github.com/catielanier/ringside/configs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var arsenalCollection *mongo.Collection = configs.GetCollection(configs.DB, "arsenalCards")
var backlashCollection *mongo.Collection = configs.GetCollection(configs.DB, "backlashCards")
var backstageCollection *mongo.Collection = configs.GetCollection(configs.DB, "backstageAreaCards")
var superstarCollection *mongo.Collection = configs.GetCollection(configs.DB, "superstars")

type cardCollection struct {
	ArsenalCards   []cards.ArsenalCard
	BacklashCards  []cards.Backlash
	BackstageCards []cards.Backstage
	Superstars     []cards.Superstar
}

func GetAllCards() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var arsenalCards []cards.ArsenalCard
		var backlashCards []cards.Backlash
		var backstageCards []cards.Backstage
		var superstarCards []cards.Superstar
		defer cancel()

		results, err := arsenalCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleCard cards.ArsenalCard
			if err := results.Decode(&singleCard); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{})
			}
			arsenalCards = append(arsenalCards, singleCard)
		}

		results, err = backlashCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleCard cards.Backlash
			if err := results.Decode(&singleCard); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{})
			}
			backlashCards = append(backlashCards, singleCard)
		}

		results, err = superstarCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleCard cards.Superstar
			if err := results.Decode(&singleCard); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{})
			}
			superstarCards = append(superstarCards, singleCard)
		}

		results, err = backstageCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		defer results.Close(ctx)

		for results.Next(ctx) {
			var singleCard cards.Backstage
			if err := results.Decode(&singleCard); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{})
			}
			backstageCards = append(backstageCards, singleCard)
		}

		allCards := cardCollection{
			ArsenalCards:   arsenalCards,
			BacklashCards:  backlashCards,
			BackstageCards: backstageCards,
			Superstars:     superstarCards,
		}
		c.JSON(http.StatusOK, allCards)
	}
}
