package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/joho/godotenv"
	"os"

	"github.com/happeens/basic-go-api/app"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "migrate":
			// migrate()
			fmt.Printf("migrating")
			return
		default:
			fmt.Printf("unknown option: %v", args[1])
		}
	}

	router := app.InitRoutes()

	for _, arg := range args {
		fmt.Printf("arg: %v\n", arg)
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	endless.ListenAndServe(port, router)
}
