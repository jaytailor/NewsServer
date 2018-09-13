package handler

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"strconv"
	"time"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {

	// Create a struct to load the videos
	videoStruct := VideoItems{Status:"OK"}

	keys, ok := r.URL.Query()["list"]
	var numberOfVideos int
	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		numberOfVideos = 10
	}else {
		items, err := strconv.Atoi(keys[0])
		if err == nil {
			fmt.Printf("Number of Videos.. ")
			fmt.Println(items)
			numberOfVideos = items
		}
	}

	videos, err := mdao.FindNumOfVideos(numberOfVideos)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	for _, element := range videos{
		videoStruct.VideoList = append(videoStruct.VideoList, element)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(videoStruct); err != nil {
		panic(err)
	}
}

func PostVideos(w http.ResponseWriter, r *http.Request) {

	var video Video
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}
	video.Id = bson.NewObjectId()
	video.PushedAt = time.Now()
	video.Video_Date = time.Now().Format("02-Jan-2006 15:04")

	if err := mdao.InsertVideos(video); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	respondWithVideoJson(w, http.StatusCreated, video)
}

func respondWithVideoJson(w http.ResponseWriter, status int, video Video){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(video); err != nil {
		panic(err)
	}
}
