package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123@/goapi?charset=utf8&loc=Local&parseTime=True")
	if err != nil {
		fmt.Printf("Error opening DB connection: %v", err)
	}
}

func Migrate(model interface{}) {
    db.AutoMigrate(model)
}

