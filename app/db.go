package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func initDb() error {
	var err error
	Db, err = gorm.Open("mysql", "root:123@/goapi?charset=utf8&loc=Local&parseTime=True")
	if err != nil {
		return err
	}

	return nil
}
