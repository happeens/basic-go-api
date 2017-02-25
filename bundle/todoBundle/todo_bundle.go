package todoBundle

import (
	"github.com/happeens/basic-go-api/app"
)

func init() {
	var todoCtrl = todoController{}
	todos := app.Router.Group("/todos")
	todos.Use(app.RequireAuth())
	{
		todos.GET("", todoCtrl.Index)
		todos.GET("/:id", todoCtrl.Show)
		todos.POST("", todoCtrl.Create)
		todos.PUT("/:id", todoCtrl.Update)
		todos.DELETE("/:id", todoCtrl.Destroy)
	}
}
