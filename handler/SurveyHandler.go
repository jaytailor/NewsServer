package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
	"time"
)


// get survey result
func GetSurveyResult(w http.ResponseWriter, r *http.Request) {

	// Create a struct to load the wish message
	surveyResult := SurveyItems{Status:"OK"}

	messages, err := mdao.FindNSurveys(1, time.Now())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	for _, element := range messages{
		surveyResult.SurveyResult = append(surveyResult.SurveyResult, element)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(surveyResult); err != nil {
		panic(err)
	}
}


// Post/save new survey
func PostSurvey(w http.ResponseWriter, r *http.Request) {

	var survey Survey
	if err := json.NewDecoder(r.Body).Decode(&survey); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}

	survey.Id = bson.NewObjectId()
	survey.StartDate = time.Now()// start now

	// put default expiry date of 2 day
	survey.EndDate = time.Now().AddDate(0, 0, 2) // end after 2 days

	if err := mdao.InsertNewSurvey(survey); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	respondWithSurveyJson(w, http.StatusCreated, survey)
}


// Update survey result from app
func UpdateSurvey(w http.ResponseWriter, r *http.Request) {

	var survey Survey
	if err := json.NewDecoder(r.Body).Decode(&survey); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}

	if err := mdao.UpdateSurvey(survey); err != nil {
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


func respondWithSurveyJson(w http.ResponseWriter, status int, survey Survey){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(survey); err != nil {
		panic(err)
	}
}