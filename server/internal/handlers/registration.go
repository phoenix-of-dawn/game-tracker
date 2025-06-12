package handlers

import "github.com/gin-gonic/gin"

func setupRegistration(router *gin.Engine) {
	router.POST("/register", registerUser)
}

func registerUser(c *gin.Context) {

}