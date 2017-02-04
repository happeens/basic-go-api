package config

import (
	"bytes"
	"gopkg.in/gin-gonic/gin.v1"

	. "github.com/happeens/basic-go-api/controllers"
)

func resource(name string, ctrl Controller, router *gin.Engine) {
	var buffer bytes.Buffer
	buffer.WriteString("/")
	buffer.WriteString(name)
	buffer.WriteString("s")
	basePath := buffer.String()

	buffer.WriteString("/:id")
	idPath := buffer.String()

	router.GET(basePath, ctrl.Index)
	router.GET(idPath, ctrl.Show)
	router.POST(basePath, ctrl.Create)
	router.PUT(idPath, ctrl.Update)
	router.DELETE(idPath, ctrl.Destroy)
}
