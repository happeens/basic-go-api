package main

import (
	"github.com/fvbock/endless"

	"github.com/happeens/basic-go-api/app"
	_ "github.com/happeens/basic-go-api/bundle/todoBundle"
	_ "github.com/happeens/basic-go-api/bundle/userBundle"
)

func main() {
	endless.ListenAndServe(app.Port, app.Router)
}
