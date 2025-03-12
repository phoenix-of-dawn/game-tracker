package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/handlers"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/igdb"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load env variables")
	}

	server_ip := os.Getenv("SERVER_IP")
	server_port := os.Getenv("SERVER_PORT")

	igdb.Setup()

	router := gin.Default()
	router.Use(handlers.CORSMiddleware())
	handlers.Setup(router)
	log.Println("Starting...")
	router.Run(server_ip + ":" + server_port)
}