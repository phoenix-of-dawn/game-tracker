package database

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/phoenix-of-dawn/game-tracker/server/api"
	"github.com/redis/go-redis/v9"
)

var ctx context.Context = context.Background()

func UserSetup() {
	_, err := Rdb.FTCreate(
		ctx,
		"idx:users",
		&redis.FTCreateOptions{
			OnJSON: true,
			Prefix: []interface{}{
				"user:",
			},
		},
		&redis.FieldSchema{
			FieldName: "$.id",
			As:        "id",
			FieldType: redis.SearchFieldTypeNumeric,
		},
		&redis.FieldSchema{
			FieldName: "$.username",
			As:        "username",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.password",
			As:        "password",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.email",
			As:        "email",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.games",
			As:        "games",
			FieldType: redis.SearchFieldTypeTag,
		},
	).Result()

	if err != nil {
		log.Fatalf("Redis FTCreate failed %s", err.Error())
	}
}

func FindUserByUsername(username string) *api.User {
	res, err := Rdb.FTSearch(ctx, "idx:users", "@username:"+username).Result()
	if err != nil {
		log.Fatalf("Could not search the Redis DB, %s", err.Error())
	}

	if res.Total == 0 {
		return nil
	} else {
		userJson, err := json.Marshal(res.Docs[0].Fields)
		if err != nil {
			log.Fatal(err.Error())
		}
		var user api.User
		json.Unmarshal(userJson, &user)
		return &user
	}
}

func UserExists(username string) bool {
	res, err := Rdb.FTSearch(ctx, "idx:users", "@username:"+username).Result()
	if err != nil {
		log.Fatalf("Could not search the Redis DB, %s", err.Error())
	}

	if res.Total == 0 {
		return false
	}
	return true
}

func GetUserByID(id int) *api.User {
	res, err := Rdb.JSONGet(ctx, "user:"+strconv.FormatInt(int64(id), 10)).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		log.Fatal(err.Error())
	}

	var user api.User
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &user
}
