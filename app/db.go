package app

import (
	"gopkg.in/mgo.v2"
)

var db *mgo.Session

func initDb() {
	var err error
	db, err = mgo.Dial("mongodb://admin:123@localhost")
	if err != nil {
		panic(err)
	}
}

func DB() *mgo.Database {
	return db.DB("godb")
}
