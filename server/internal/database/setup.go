package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/phoenix-of-dawn/game-tracker/server/internal/user"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Client *mongo.Client
var coll *mongo.Collection

func Setup() {
	url := os.Getenv("DATABASE_URL")
	username := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	log.Print(url, username, pass)
	Client, _ = mongo.Connect(options.Client().ApplyURI("mongodb://" + username + ":" + pass + "@" + url))
	ctx := context.Background()

	defer func() {
		if err := Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	print(Client)
	err := Client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Panic(err)
	}

	coll = Client.Database("test").Collection("users")

	user.Setup(Client)

	// Will throw an error if the definitions of the index models change
	createIndexes()
}

func createIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := coll.Indexes().CreateMany(ctx, indexModels)

	if err != nil {
		log.Panic(err)
	}
}
