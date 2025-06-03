package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Client *mongo.Client

func Setup() {
	url := os.Getenv("DATABASE_URL")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	Client, _ = mongo.Connect(options.Client().ApplyURI("mongodb://" + user + ":" + pass + "@" + url))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel();

	defer func() {
		if err := Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()


	_ = Client.Ping(ctx, readpref.Primary());
}
