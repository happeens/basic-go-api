package todoBundle

import (
	"gopkg.in/mgo.v2/bson"
)

type todo struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Description string
	Done        bool
}
