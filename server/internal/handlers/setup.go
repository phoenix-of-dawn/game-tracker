package handlers

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	setupGames(router)
}
