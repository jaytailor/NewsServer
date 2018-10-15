package handler

import (
"encoding/json"
"net/http"
"gopkg.in/mgo.v2/bson"
"strconv"
	"time"
)

func GetAllEditorial(w http.ResponseWriter, r *http.Request) {

	// Create a struct to load the editorial articles
	editorialStruct := EditorialItem{Status:"OK"}

	keys, ok := r.URL.Query()["list"]
	var numberOfArticles int
	if !ok || len(keys[0]) < 1 {
		numberOfArticles = 10
	}else {
		items, err := strconv.Atoi(keys[0])
		if err == nil {
			numberOfArticles = items
		}
	}

	articles, err := mdao.FindNumOfEditorial(numberOfArticles)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	for _, element := range articles{
		editorialStruct.ArticleList = append(editorialStruct.ArticleList, element)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(editorialStruct); err != nil {
		panic(err)
	}
}

func PostEditorial(w http.ResponseWriter, r *http.Request) {

	var article Editorial
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}

	article.Id = bson.NewObjectId()
	article.PushedAt = time.Now()
	article.PublishedAt = time.Now().Format("02-Jan-2006 15:04")

	if err := mdao.InsertEditorial(article); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	respondWithEditorialJson(w, http.StatusCreated, article)
}

func respondWithEditorialJson(w http.ResponseWriter, status int, article Editorial){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(article); err != nil {
		panic(err)
	}
}
