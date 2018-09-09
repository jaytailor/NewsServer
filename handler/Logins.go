package handler

import (
	"gopkg.in/mgo.v2/bson"
)

type Logins struct {
	Id	bson.ObjectId	`bson:"_id" json:"id"`
	User string `bson:"user" json:"user"`
	Password string `bson:"password" json:"password"`
	Role string `bson:"role" json:"role"`
}

type Success struct {
	Authenticated string `bson:"Authenticated" json:"Authenticated"`
	Message string `bson:"message" json:"message"`
}

type Credentials []Logins
