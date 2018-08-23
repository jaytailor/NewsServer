package handler

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type VideoItems struct {
	Status string `json:"status" bson:"status"`
	VideoList []Video  `json:"video_list" bson:"video_list"`
}

type Video struct {
	Id	bson.ObjectId	`bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Url string `bson:"url" json:"url"`
	Video_Date time.Time `bson:"video_date" json:"video_date"`
}

type VideoList []Video

