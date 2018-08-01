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
	Expired bool `bson:"expired" json:"expired"`
	StartDate time.Time `bson:"startdate" json:"startdate"`
	EndDate time.Time	`bson:"enddate" json:"enddate"`
	ImpressionLimit string `bson:"impressionlimit" json:"impressionlimit"`
	ImpressionFreq string `bson:"impressionfreq" json:"impressionfreq"`
	Priority string `bson:"priority" json:"priority"`
	CurrentImpressionCount string `bson:"currentimpressioncount" json:"currentimpressioncount"`
}

type AdsList []Campaigns
