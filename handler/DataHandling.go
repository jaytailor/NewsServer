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
	VIDEOS = "videos"
	EDITORIALS = "editorials"
)

func (m *DbDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)

}


// Helper mongo functions for News

func (m *DbDAO) FindAllNews() ([]NewsModel, error) {
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

func (m *DbDAO) InsertNews(news NewsModel) error {
	err := db.C(COLLECTION).Insert(&news)
	return err
}

func (m *DbDAO) DeleteNews(news NewsModel) error {
	err := db.C(COLLECTION).Remove(&news)
	return err
}

func (m *DbDAO) UpdateNews(news NewsModel) error {
	err := db.C(COLLECTION).UpdateId(news.Id, &news)
	return err
}

// Helper mongo functions for ads

func (m *DbDAO) InsertAds(ads Campaigns) error {
	err := db.C(ADSTABLE).Insert(&ads)
	return err
}

func (m *DbDAO) FindAllCampaigns() ([]Campaigns, error) {
	var campaigns []Campaigns
	err := db.C(ADSTABLE).Find(bson.M{}).All(&campaigns)
	return campaigns, err
}

func (m *DbDAO) FindNumOfAds(number int) ([]Campaigns, error) {
	var campaigns []Campaigns
	err := db.C(ADSTABLE).Find(bson.M{}).Sort("-start_date").Limit(number).All(&campaigns)
	return campaigns, err
}

func (m *DbDAO) DeleteAds(campaign Campaigns) error {
	err := db.C(ADSTABLE).Remove(&campaign)
	return err
}

func (m *DbDAO) DeleteAdById(id string) error {
	err := db.C(ADSTABLE).Remove(bson.M{"id": id})
	return err
}

// Helper mongo functions for video

func (m *DbDAO) InsertVideos(video Video) error {
	err := db.C(VIDEOS).Insert(&video)
	return err
}

func (m *DbDAO) FindAllVideos() ([]Video, error) {
	var video []Video
	err := db.C(VIDEOS).Find(bson.M{}).All(&video)
	return video, err
}

func (m *DbDAO) FindNumOfVideos(number int) ([]Video, error) {
	var video []Video
	err := db.C(VIDEOS).Find(bson.M{}).Sort("-video_date").Limit(number).All(&video)
	return video, err
}

// Helper mongo functions for editorial

func (m *DbDAO) InsertEditorial(editorial Editorial) error {
	err := db.C(EDITORIALS).Insert(&editorial)
	return err
}

func (m *DbDAO) FindAllEditorial() ([]Editorial, error) {
	var editorials []Editorial
	err := db.C(EDITORIALS).Find(bson.M{}).All(&editorials)
	return editorials, err
}

func (m *DbDAO) FindNumOfEditorial(number int) ([]Editorial, error) {
	var editorials []Editorial
	err := db.C(EDITORIALS).Find(bson.M{}).Sort("-published_at").Limit(number).All(&editorials)
	return editorials, err
}