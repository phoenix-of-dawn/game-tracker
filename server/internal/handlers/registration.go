package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/user"
)

func setupRegistration(router *gin.Engine) {
	router.POST("/register", registerUser)
}

func registerUser(c *gin.Context) {
	newUserRequest := &user.UserRegisterRequest{}

	if err := c.BindJSON(newUserRequest); err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusBadRequest, newUserRequest)
		return
	}

	_, err := user.RegisterUser(newUserRequest)
	if err != nil {
		log.Panic(err)
	}

	c.IndentedJSON(http.StatusAccepted, nil)
}
