package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strconv"
)


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Begun Darshan Server, %q", html.EscapeString(r.URL.Path))
}

func GetAllNews(w http.ResponseWriter, r *http.Request) {

	mainNewsStruct := NewsItem{Status:"OK"}

	keys, ok := r.URL.Query()["list"]
	var numberOfNews int
	if !ok || len(keys[0]) < 1 {
		//fmt.Println("Url Param 'key' is missing")
		numberOfNews = 10
	}else {
		items, err := strconv.Atoi(keys[0])
		if err == nil {
			//fmt.Printf("Number of news.. ")
			numberOfNews = items
		}
	}

	samachar, err := mdao.FindNumOfNews(numberOfNews)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	for _, element := range samachar{
		mainNewsStruct.NewsItems = append(mainNewsStruct.NewsItems, element)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(mainNewsStruct); err != nil {
		panic(err)
	}
}

func PostNews(w http.ResponseWriter, r *http.Request) {

	var samachar NewsModel
	if err := json.NewDecoder(r.Body).Decode(&samachar); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}
	samachar.Id = bson.NewObjectId()
	samachar.PublishedAt = time.Now()

	if err := mdao.InsertNews(samachar); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	respondWithJson(w, http.StatusCreated, samachar)
}

func FindSpecificNews(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var samachar NewsModel
	samachar, err := mdao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		panic(err)
		return
	}
	respondWithJson(w, http.StatusOK, samachar)
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {

	var samachar NewsModel
	if err := json.NewDecoder(r.Body).Decode(&samachar); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}
	if err := mdao.UpdateNews(samachar); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{"result": "success"}); err != nil {
		panic(err)
	}
}

func DeleteNews(w http.ResponseWriter, r *http.Request) {

	var samachar NewsModel
	if err := json.NewDecoder(r.Body).Decode(&samachar); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}
	if err := mdao.DeleteNews(samachar); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{"result": "success"}); err != nil {
		panic(err)
	}
}

func respondWithJson(w http.ResponseWriter, status int, samachar NewsModel){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(samachar); err != nil {
		panic(err)
	}
}


