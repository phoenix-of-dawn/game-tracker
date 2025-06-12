package user

import (
	"errors"
	"log"

	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
)

var (
	node, _ = snowflake.NewNode(1)
)

func RegisterUser(request UserRegisterRequest) (*User, error) {
	user, err := makeAndValidateUser(request)

	if err != nil {
		return nil, err
	}

	return InsertUser(user)
}

func makeAndValidateUser(request UserRegisterRequest) (*User, error) {
	email := request.Email
	username := request.Username
	password := request.Password

	if len(password) < 6 {
		return nil, errors.New("invalid password length")
	}

	if UserExists(email) {
		return nil, errors.New("user already exists")
	}

	password, err := hashPassword(password)

	if err != nil {
		log.Panic(err)
	}

	user := User{
		Id:       node.Generate(),
		Email:    email,
		Username: username,
		Password: password,
		Games:    []string{},
	}

	return &user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
