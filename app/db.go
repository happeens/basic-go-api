package app

import (
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func initDb() {
	// get database config
	host := Env("DB_HOST")
	port := Env("DB_PORT")
	user := Env("DB_USERNAME")
	pass := Env("DB_PASSWORD")
	name := Env("DB_DATABASE")

	dialString := "mongodb://" + user + ":" + pass + "@" + host + ":" + port + "/" + name
	Log.Debugf("dialstring: %v", dialString)

	con, err := mgo.Dial(dialString)
	if err != nil {
		panic(err)
	}

	db = con.DB(Env("DB_DATABASE"))
}

func DB() *mgo.Database {
	return db
}
