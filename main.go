package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/happeens/basic-go-api/config"
)

func main() {
    args := os.Args
    if(len(args) > 1) {
        switch(args[1]) {
        case "migrate":
            migrate()
            return
        default:
            fmt.Printf("unknown option: %v", args[1])
        }
    }

	router := config.InitRoutes()

    for _, arg := range args {
        fmt.Printf("arg: %v\n", arg)
    }

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	router.Run(port)
}

