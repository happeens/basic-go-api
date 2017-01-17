package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"api/models"
)

func main() {
	db, err := gorm.Open("mysql", "gouser:w00f1fad@/todos")
	if err != nil {
		fmt.Printf("Error opening DB connection: %v", err)
	}

	defer db.Close()

	db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&models.User{})
}
