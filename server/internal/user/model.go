package user

import "github.com/bwmarrin/snowflake"

type User struct {
	Id       snowflake.ID `json:"_id" bson:"_id"`
	Username string       `json:"username" bson:"username"`
	Password string       `json:"password" bson:"password"`
	Email    string       `json:"email" bson:"email"`
	Games    []string     `json:"games" bson:"games"`
}

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
