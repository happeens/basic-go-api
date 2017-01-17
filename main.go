package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"api/config"
)

func main() {
	router := config.InitRoutes()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	router.Run(port)
}
