package app

import (
	"github.com/joho/godotenv"
	"os"
)

var Port = ":8080"
var Env = "dev"

func initConf() {
	godotenv.Load(".env")

	portEnv := os.Getenv("PORT")
	if portEnv != "" {
		//TODO: check valid port
		Port = ":" + portEnv
	}

	modeEnv := os.Getenv("ENV")
	if modeEnv != "" {
		//TODO: check valid mode
		Env = modeEnv
	}
}
