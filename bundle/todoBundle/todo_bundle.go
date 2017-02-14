package todoBundle

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/happeens/basic-go-api/app"
)

func init() {
	var todoCtrl = todoController{}
	todos := app.Router.Group("/todos")
	{
		todos.GET("", todoCtrl.Index)
		todos.GET("/:id", todoCtrl.Show)
		todos.POST("", todoCtrl.Create)
		todos.PUT("/:id", todoCtrl.Update)
		todos.DELETE("/:id", todoCtrl.Destroy)
	}
}

func initRoutes(r *gin.Engine) {
}
