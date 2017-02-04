package config

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/happeens/basic-go-api/controllers"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	resource("todo", controllers.TodoController{}, router)

	return router
}
