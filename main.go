package main

import (
	"log"
	"os"

	"github.com/arifmr/dbo/database"
	"github.com/arifmr/dbo/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal("Failed to load environment variables:", err)
	}

	err = database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	router := gin.Default()
	routes.SetupRoutes(router)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}

	router.Run(port)
}

func loadEnv() error {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		log.Println("No .env file found, reading environment variables from Docker")
	}

	return nil
}
