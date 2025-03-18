package database

import (
	"context"
	"log"

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
			As: "id",
			FieldType: redis.SearchFieldTypeNumeric,
		},
		&redis.FieldSchema{
			FieldName: "$.username",
			As: "username",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.password",
			As: "password",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.email",
			As: "email",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.games",
			As: "games",
			FieldType: redis.SearchFieldTypeTag,
		},
	).Result()

	if err != nil {
		log.Fatalf("Redis FTCreate failed %s", err.Error())
	}
}