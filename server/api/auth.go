package api

import "github.com/bwmarrin/snowflake"

type User struct {
	Id       snowflake.ID `json:"id"`
	Username string       `json:"username"`
	Password string       `json:"password"`
	Email    string       `json:"email"`
}
