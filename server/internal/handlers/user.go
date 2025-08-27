package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/user"
)

func setupUser(router *gin.Engine) {
	router.GET("/user", getUser)
}

func getUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	userReq, err := user.GetUserByID(int(id))

	if err != nil {
		log.Panic(err)
	}

	userResponse := user.UserRequest{
		Id:       userReq.Id,
		Username: userReq.Username,
		Games:    userReq.Games,
	}

	c.IndentedJSON(http.StatusOK, userResponse)
}
