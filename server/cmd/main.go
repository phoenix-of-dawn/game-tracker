package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/handlers"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/igdb"
)

func main() {
	igdb.Setup()

	router := gin.Default()
	handlers.Setup(router)
	log.Println("Starting...")
	router.Run("localhost:8080")
}