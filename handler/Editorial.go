package handler

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type EditorialItem struct {
	Status string `json:"status" bson:"status"`
	ArticleList []Editorial  `json:"article_list" bson:"article_list"`
}

type Editorial struct {
	Id	bson.ObjectId	`bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
	Image []string `bson:"image" json:"image"`
	Writer string `bson:"writer" json:"writer"`
	PublishedAt string	`bson:"published_at" json:"published_at"`
	PushedAt time.Time `bson:"pushed_at" json:"pushed_at"`
	Status string `bson:"status" json:"status"`
	Extra string `bson:"extra" json:"extra"`
}

type ArticleList []Editorial
