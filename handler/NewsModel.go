package handler

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	)

type NewsItem struct {
	Status string `json:"status" bson:"status"`
	NewsItems []NewsModel  `json:"newsitems" bson:"newsitems"`
}

type NewsModel struct {
	Id	bson.ObjectId	`bson:"_id" json:"id"`
	Writer string `bson:"writer" json:"writer"`
	Title string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
	Image string `bson:"image" json:"image"`
	PublishedAt string	`bson:"published_at" json:"published_at"`
	PushedAt time.Time `bson:"pushed_at" json:"pushed_at"`
	IsBreaking bool `bson:"is_breaking" json:"is_breaking"`
	Status string `bson:"status" json:"status"`
	Extra string `bson:"extra" json:"extra"`
}

type NewsList []NewsModel

