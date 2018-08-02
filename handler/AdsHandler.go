package handler

import (
"encoding/json"
"net/http"

"gopkg.in/mgo.v2/bson"
	"fmt"
	"strconv"
)

type ID struct {
	id string
}

func GetAllAds(w http.ResponseWriter, r *http.Request) {
	mdao := DbDAO{Server:"localhost", Database:"news"}
	mdao.Connect()

	mainNewsStruct := AdsItems{Status:"OK"}

	keys, ok := r.URL.Query()["list"]
	var numberOfAds int
	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		numberOfAds = 10
	}else {
		items, err := strconv.Atoi(keys[0])
		if err == nil {
			fmt.Printf("Number of Ads.. ")
			fmt.Println(items)
			numberOfAds = items
		}
	}

	ads, err := mdao.FindNumOfAds(numberOfAds)

	if err != nil {
		//respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	for _, element := range ads{
		mainNewsStruct.AdsItem = append(mainNewsStruct.AdsItem, element)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(mainNewsStruct); err != nil {
		panic(err)
	}
}

func PostAds(w http.ResponseWriter, r *http.Request) {
	mdao := DbDAO{Server:"localhost", Database:"news"}
	mdao.Connect()
	//defer session.Close()

	defer r.Body.Close()
	var ads Campaigns
	if err := json.NewDecoder(r.Body).Decode(&ads); err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}
	ads.Id = bson.NewObjectId()

	if err := mdao.InsertAds(ads); err != nil {
		//respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	respondWithAdsJson(w, http.StatusCreated, ads)
}

func DeleteAds(w http.ResponseWriter, r *http.Request) {

	mdao := DbDAO{Server:"localhost", Database:"news"}
	mdao.Connect()

	defer r.Body.Close()
	var ad Campaigns

	if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		fmt.Println("error in decoding")
		panic(err)
		return
	}

	if err := mdao.DeleteAds(ad); err != nil {
		//respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{"result": "success"}); err != nil {
		panic(err)
	}
}

func respondWithAdsJson(w http.ResponseWriter, status int, ads Campaigns){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(ads); err != nil {
		panic(err)
	}
}

