package handler

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type AdsItems struct {
	Status string `json:"status" bson:"status"`
	AdsItem []Campaigns `json:"campaigns" bson:"campaigns"`
}

type Campaigns struct {
	Id	bson.ObjectId	`bson:"_id" json:"id"`
	ImageUrl string `bson:"imageurl" json:"imageurl"`
	StartDate time.Time `bson:"start_date" json:"start_date"`
	EndDate time.Time	`bson:"end_date" json:"end_date"`
	ImpressionLimit string `bson:"impression_limit" json:"impression_limit"`
	ImpressionFreq string `bson:"impression_freq" json:"impression_freq"`
	Priority int `bson:"priority" json:"priority"`
	CurrentImpressionCount string `bson:"current_impression_count" json:"current_impression_count"`
	Adtype string `bson:"adtype" json:"adtype"`
	Status string `bson:"status" json:"status"`
	Extra string `bson:"extra" json:"extra"`
}

type AdsList []Campaigns
