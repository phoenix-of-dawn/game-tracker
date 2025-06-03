package api

import "github.com/bwmarrin/snowflake"

type User struct {
	Id       snowflake.ID `json:"id" bson:"id"`
	Username string       `json:"username" bson:"username"`
	Password string       `json:"password" bson:"password"`
	Email    string       `json:"email" bson:"email"`
	Games    []string     `json:"games" bson:"games"`
}
