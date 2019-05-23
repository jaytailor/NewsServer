package handler

import (
"encoding/json"
"net/http"

"gopkg.in/mgo.v2/bson"
"strconv"
	"time"
)


// get number of congratulatory wish messages. 
func GetAllWishes(w http.ResponseWriter, r *http.Request) {

	// Create a struct to load the wish message
	wishMessages := WishItems{Status:"OK"}

	keys, ok := r.URL.Query()["list"]
	var numOfMessages int
	if !ok || len(keys[0]) < 1 {
		numOfMessages = 5 // by default return 5 messages 
	}else {
		items, err := strconv.Atoi(keys[0])
		if err == nil {
			numOfMessages = items
		}
	}

	messages, err := mdao.FindNMessages(numOfMessages, time.Now())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	for _, element := range messages{
		wishMessages.WishItems = append(wishMessages.WishItems, element)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(wishMessages); err != nil {
		panic(err)
	}
}


// Post/save congratulatory messages. 
func PostWishMessage(w http.ResponseWriter, r *http.Request) {

	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}

	message.Id = bson.NewObjectId()
	message.StartDate = time.Now()// start now

	// put default expiry date of 3 days
	message.EndDate = time.Now().AddDate(0, 0, 3) // end after three days

	if err := mdao.InsertMessages(message); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	respondWithMessageJson(w, http.StatusCreated, message)
}


func respondWithMessageJson(w http.ResponseWriter, status int, message Message){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}
