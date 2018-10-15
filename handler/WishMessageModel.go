package handler

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type WishItems struct {
	Status string `json:"status" bson:"status"`
	WishItems []Message `json:"messages" bson:"messages"`
}

type Message struct {
	Id	bson.ObjectId	`bson:"_id" json:"id"`
	ImageUrl string `bson:"imageurl" json:"imageurl"`
	StartDate time.Time `bson:"start_date" json:"start_date"`
	EndDate time.Time	`bson:"end_date" json:"end_date"`
	MessageText string `bson:"message_text" json:"message_text"`
	Extra string `bson:"extra" json:"extra"`
}

type MessageList []Message
