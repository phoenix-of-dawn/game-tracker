package api

import "github.com/bwmarrin/snowflake"

type User struct {
	Id snowflake.ID `redis:"id"`
	Username string `redis:"username"`
	Password string `redis:"password"`
	Email string `redis:"email"`
}