package configs

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func ConnectDB() *mongo.Client {
	err := godotenv.Load()
	mongoDBURI := os.Getenv("MONGODB_URI")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongo")
	return client
}
