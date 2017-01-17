package config

import (
	"gopkg.in/gin-gonic/gin.v1"

	"api/controllers"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	resource("todo", controllers.TodoController{}, router)

	return router
}
