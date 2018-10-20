package handler

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
)

type DbDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	DATABASE = "begundarshan"
	NEWS_TABLE = "news"
	ADS_TABLE = "ads"
	VIDEO_TABLE = "videos"
	EDITORIALS = "editorials"
	LOGIN = "login"
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
	err := db.C(NEWS_TABLE).Find(bson.M{}).All(&news)
	return news, err
}

func (m *DbDAO) FindNumOfNews(number int) ([]NewsModel, error) {
	var news []NewsModel
	err := db.C(NEWS_TABLE).Find(bson.M{}).Sort("-pushed_at").Limit(number).All(&news)
	return news, err
}

func (m *DbDAO) FindById(id string) (NewsModel, error) {
	var news NewsModel
	err := db.C(NEWS_TABLE).FindId(bson.ObjectIdHex(id)).One(&news)
	return news, err
}

func (m *DbDAO) InsertNews(news NewsModel) error {
	err := db.C(NEWS_TABLE).Insert(&news)
	return err
}

func (m *DbDAO) DeleteNews(news NewsModel) error {
	err := db.C(NEWS_TABLE).Remove(&news)
	return err
}

func (m *DbDAO) UpdateNews(news NewsModel) error {
	err := db.C(NEWS_TABLE).UpdateId(news.Id, &news)
	return err
}

// Helper mongo functions for ads

func (m *DbDAO) InsertAds(ads Campaigns) error {
	err := db.C(ADS_TABLE).Insert(&ads)
	return err
}

func (m *DbDAO) FindAllCampaigns() ([]Campaigns, error) {
	var campaigns []Campaigns
	err := db.C(ADS_TABLE).Find(bson.M{}).All(&campaigns)
	return campaigns, err
}

// function to return active and unexpired campaign only
func (m *DbDAO) FindNumOfAds(number int, nowdate time.Time) ([]Campaigns, error) {
	var campaigns []Campaigns
	err := db.C(ADS_TABLE).Find(bson.M{"end_date":bson.M{"$gt":nowdate}, "status":"active"}).Sort( "priority").Limit(number).All(&campaigns)
	return campaigns, err
}

func (m *DbDAO) FindAdsOfPriority(priority int, nowdate time.Time) ([]Campaigns, error) {
	var campaigns []Campaigns
	fmt.Println()
	err := db.C(ADS_TABLE).Find(bson.M{"end_date":bson.M{"$gt":nowdate},"priority":priority, "status":"active"}).Sort("-start_date").All(&campaigns)
	return campaigns, err
}

func (m *DbDAO) FindAdsAbovePriority(priority int, nowdate time.Time) ([]Campaigns, error) {
	var campaigns []Campaigns
	fmt.Println()
	err := db.C(ADS_TABLE).Find(bson.M{"end_date":bson.M{"$gt":nowdate},"priority":bson.M{"$gt":priority}, "status":"active"}).Sort( "priority").All(&campaigns)
	return campaigns, err
}

func (m *DbDAO) DeleteAds(campaign Campaigns) error {
	err := db.C(ADS_TABLE).Remove(&campaign)
	return err
}

func (m *DbDAO) DeleteAdById(id string) error {
	err := db.C(ADS_TABLE).Remove(bson.M{"id": id})
	return err
}

// Helper mongo functions for video

func (m *DbDAO) InsertVideos(video Video) error {
	err := db.C(VIDEO_TABLE).Insert(&video)
	return err
}

func (m *DbDAO) FindAllVideos() ([]Video, error) {
	var video []Video
	err := db.C(VIDEO_TABLE).Find(bson.M{}).All(&video)
	return video, err
}

func (m *DbDAO) FindNumOfVideos(number int) ([]Video, error) {
	var video []Video
	err := db.C(VIDEO_TABLE).Find(bson.M{}).Sort("-pushed_at").Limit(number).All(&video)
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
	err := db.C(EDITORIALS).Find(bson.M{}).Sort("-pushed_at").Limit(number).All(&editorials)
	return editorials, err
}

// Helper mongo functions for user logins

func (m *DbDAO) CreateLogins(credentials Logins) error {
	err := db.C(LOGIN).Insert(&credentials)
	return err
}

func (m *DbDAO) FindUsers(username string) ([]Logins, error) {
	var credentials []Logins
	err := db.C(LOGIN).Find(bson.M{"username": username}).All(&credentials)
	return credentials, err
}