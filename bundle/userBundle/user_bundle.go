package userBundle

import (
	"github.com/happeens/basic-go-api/app"
)

func init() {
	var userCtrl = userController{}

	app.Router.POST("/authenticate", userCtrl.Authenticate)

	users := app.Router.Group("/users")
	{
		users.GET("", userCtrl.Index)
		users.GET("/:id", userCtrl.Show)
		users.POST("", userCtrl.Create)
	}
}
