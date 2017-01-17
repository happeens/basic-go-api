package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type Controller interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}
