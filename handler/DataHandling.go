package handler

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NewsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "news"
)

func (m *NewsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)

}

func (m *NewsDAO) FindAll() ([]NewsModel, error) {
	var news []NewsModel
	err := db.C(COLLECTION).Find(bson.M{}).All(&news)
	return news, err
}

func (m *NewsDAO) FindNumOfNews(number int) ([]NewsModel, error) {
	var news []NewsModel
	err := db.C(COLLECTION).Find(bson.M{}).Sort("-published_at").Limit(number).All(&news)
	return news, err
}

func (m *NewsDAO) FindById(id string) (NewsModel, error) {
	var news NewsModel
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&news)
	return news, err
}

func (m *NewsDAO) Insert(news NewsModel) error {
	err := db.C(COLLECTION).Insert(&news)
	return err
}

func (m *NewsDAO) Delete(news NewsModel) error {
	err := db.C(COLLECTION).Remove(&news)
	return err
}

func (m *NewsDAO) Update(news NewsModel) error {
	err := db.C(COLLECTION).UpdateId(news.Id, &news)
	return err
}

