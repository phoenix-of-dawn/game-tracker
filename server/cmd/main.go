package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/database"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/handlers"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/igdb"
)

func main() {
	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load env variables")
	}

	server_ip := os.Getenv("SERVER_IP")
	server_port := os.Getenv("SERVER_PORT")

	// Set up the API
	igdb.Setup()

	// Set up the db
	database.Setup()

	defer func() {
		if err := database.Client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Make router
	router := gin.Default()
	router.Use(handlers.CORSMiddleware())
	router.SetTrustedProxies(nil)

	// Set up the routes
	handlers.Setup(router)

	router.Run(server_ip + ":" + server_port)
}
