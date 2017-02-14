package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func initDb() error {
	var err error
	db, err = gorm.Open("mysql", "root:123@/goapi?charset=utf8&loc=Local&parseTime=True")
	if err != nil {
		return err
	}

	return nil
}

func DB() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = gorm.Open("mysql", "root:123@/goapi?charset=utf8&loc=Local&parseTime=True")
	if err != nil {
		panic(err)
	}

	return db
}
