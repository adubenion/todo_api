package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Description string `bson:"description" json:"description"`
	Completed bool `bson:"completed" json:"completed"`
	User bson.ObjectId `bson:"_id" json:"_id"`
}
