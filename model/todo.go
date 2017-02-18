package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Description string
	Done        bool
}
