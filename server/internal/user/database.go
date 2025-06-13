package user

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var coll *mongo.Collection

func Setup(client *mongo.Client) {
	coll = client.Database("test").Collection("users")
}

func InsertUser(user *User) (*User, error) {
	_, err := coll.InsertOne(context.Background(), user)
	return user, err
}

func GetUserByEmail(email string) (*User, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var result User

	user, err := getUserWithFilter(filter, result)

	return user, err
}

func UserExists(email string) bool {
	user, _ := GetUserByEmail(email)

	return user != nil
}

func GetUserByID(id int) (*User, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var result User
	user, err := getUserWithFilter(filter, result)

	return user, err
}

func getUserWithFilter(filter bson.D, result User) (*User, error) {
	err := coll.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		log.Panic(err)
	}

	return &result, nil
}
