package database

import (
	"context"
	"log"

	"github.com/phoenix-of-dawn/game-tracker/server/api"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var coll *mongo.Collection = client.Database("main").Collection("users")

func GetUserByEmail(email string) (*api.User, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var result api.User
	
	user, err := getUserWithFilter(filter, result)

	return user, err;
}

func UserExists(email string) bool {
	user, _ := GetUserByEmail(email)

	return user != nil
}

func GetUserByID(id int) (*api.User, error) {
	filter := bson.D{{Key: "id", Value: id}}

	var result api.User
	user, err := getUserWithFilter(filter, result)

	return user, err
}

func getUserWithFilter(filter bson.D, result api.User) (*api.User, error) {
	err := coll.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		log.Panic(err)
	}

	return &result, nil
}
