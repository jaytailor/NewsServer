package handler

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DbDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "news"
	ADSTABLE = "ads"
)

func (m *DbDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)

}

func (m *DbDAO) FindAll() ([]NewsModel, error) {
	var news []NewsModel
	err := db.C(COLLECTION).Find(bson.M{}).All(&news)
	return news, err
}

func (m *DbDAO) FindNumOfNews(number int) ([]NewsModel, error) {
	var news []NewsModel
	err := db.C(COLLECTION).Find(bson.M{}).Sort("-published_at").Limit(number).All(&news)
	return news, err
}

func (m *DbDAO) FindById(id string) (NewsModel, error) {
	var news NewsModel
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&news)
	return news, err
}

func (m *DbDAO) Insert(news NewsModel) error {
	err := db.C(COLLECTION).Insert(&news)
	return err
}

func (m *DbDAO) Delete(news NewsModel) error {
	err := db.C(COLLECTION).Remove(&news)
	return err
}

func (m *DbDAO) Update(news NewsModel) error {
	err := db.C(COLLECTION).UpdateId(news.Id, &news)
	return err
}


func (m *DbDAO) InsertAds(ads Campaigns) error {
	err := db.C(ADSTABLE).Insert(&ads)
	return err
}

func (m *DbDAO) FindAllCampaigns() ([]Campaigns, error) {
	var campaigns []Campaigns
	err := db.C(ADSTABLE).Find(bson.M{}).All(&campaigns)
	return campaigns, err
}

