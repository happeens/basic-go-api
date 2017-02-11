package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/happeens/gin-jwt"
	"github.com/joho/godotenv"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/happeens/basic-go-api/controllers"
)

var Port string
var Router *gin.Engine

var authMW *jwt.GinJWTMiddleware

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port = fmt.Sprintf(":%v", os.Getenv("PORT"))
	Router = gin.Default()

	setupAuth()
	setupRoutes()
}

func setupRoutes() {
	resource("todo", controllers.TodoController{})
}

func setupAuth() {
	authMW = &jwt.GinJWTMiddleware{
		Realm: "test zone",
		//TODO: secret from env, secret generator
		Key:        []byte("secret"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(userID string, password string, c *gin.Context) (string, bool) {
			return "admin", true
		},
		Authorizator: func(userID string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:Authorization",
	}

	Router.POST("/login", authMW.LoginHandler)
}
