package config

import (
	"fmt"

	. "github.com/happeens/basic-go-api/controllers"
)

func resource(name string, ctrl Controller, methods ...string) {
	if len(methods) <= 0 {
		methods = append(methods, []string{"index", "show", "create", "update", "destroy"}...)
	}

	groupName := "/" + name + "s"
	g := Router.Group(groupName)
	g.Use(authMW.MiddlewareFunc())

	for _, method := range methods {
		switch method {
		case "index":
			g.GET("", ctrl.Index)
		case "show":
			g.GET("/:id", ctrl.Show)
		case "create":
			g.POST("", ctrl.Create)
		case "update":
			g.PUT("/:id", ctrl.Update)
		case "destroy":
			g.DELETE("/:id", ctrl.Destroy)
		default:
			fmt.Printf("Unrecognized resource method: %v", method)
		}
	}
}
