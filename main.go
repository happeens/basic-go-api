package main

import (
	"github.com/fvbock/endless"

	"api/config"
)

func main() {
	endless.ListenAndServe(config.Port, config.Router)
}
