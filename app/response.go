package app

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func Respond(c *gin.Context, code int, data ...interface{}) {
	if err, ok := data[0].(error); ok {
		c.JSON(code, gin.H{"error": err.Error()})
	} else {
		c.JSON(code, gin.H{"data": data})
	}
}
