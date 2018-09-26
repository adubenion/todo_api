package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Email string `bson:"email" json:"password"`
	Username string `bson:"username" json:"password"`
	Password string `bson:"password" json:"password"`
}
