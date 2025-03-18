package database

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb redis.Client

func Setup() {
	Rdb = *redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DATABASE_URL"),
		Password: "",
		DB:       0,
	})
}
