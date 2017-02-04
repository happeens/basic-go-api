package main

import (
    "fmt"

	"github.com/happeens/basic-go-api/models"
)

func migrate() {
    fmt.Printf("migrating...\n");

    var todoModel = models.Todo{}
    models.Migrate(&todoModel)
}

