package handler

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

func GetNumBNews(w http.ResponseWriter, r *http.Request) {

	mainBNewsStruct := BreakingNewsItem{Status:"OK"}

	keys, ok := r.URL.Query()["list"]
	var numberOfNews int
	if !ok || len(keys[0]) < 1 {
		numberOfNews = 10
	}else {
		items, err := strconv.Atoi(keys[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		numberOfNews = items
	}

	breaking, err := mdao.FindNumOfBNews(numberOfNews)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, element := range breaking{
		mainBNewsStruct.BNewsItems = append(mainBNewsStruct.BNewsItems, element)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(mainBNewsStruct); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func PostBreakingNews(w http.ResponseWriter, r *http.Request) {

	var breaking BreakingNews
	if err := json.NewDecoder(r.Body).Decode(&breaking); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	breaking.Id = bson.NewObjectId()
	loc, _ := time.LoadLocation("Asia/Kolkata")
	breaking.PushedAt = time.Now().In(loc)
	breaking.PublishedAt = time.Now().In(loc).Format("02-Jan-2006 15:04")

	if err := mdao.InsertBNews(breaking); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithBNewsJson(w, http.StatusCreated, breaking)
}

func respondWithBNewsJson(w http.ResponseWriter, status int, samachar BreakingNews){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(samachar); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}



