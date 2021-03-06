package main

import (
	"github.com/fvbock/endless"

	"github.com/happeens/basic-go-api/app"
	_ "github.com/happeens/basic-go-api/bundle/todoBundle"
	_ "github.com/happeens/basic-go-api/bundle/userBundle"
)

func main() {
	port := ":" + app.Env("PORT", "8000")
	endless.ListenAndServe(port, app.Router)
}
