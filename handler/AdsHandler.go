package handler

import (
"encoding/json"
"net/http"

"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
	"math/rand"
)

type ID struct {
	id string
}


func GetAllAds(w http.ResponseWriter, r *http.Request) {

	mainNewsStruct := AdsItems{Status:"OK"}

	keys, ok := r.URL.Query()["list"]
	var numberOfAds int
	if !ok || len(keys[0]) < 1 {
		numberOfAds = 10
	}else {
		items, err := strconv.Atoi(keys[0])
		if err == nil {
			numberOfAds = items
		}
	}

	ads, err := mdao.FindNumOfAds(numberOfAds, time.Now())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	for _, element := range ads{

		// only add campaign of priority of 10 above
		if element.Priority >= 10{
			mainNewsStruct.AdsItem = append(mainNewsStruct.AdsItem, element)
		}
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(mainNewsStruct); err != nil {
		panic(err)
	}
}

func GetAdsWithPriority(w http.ResponseWriter, r *http.Request) {

	mainNewsStruct := AdsItems{Status:"OK"}

	priority, ok := r.URL.Query()["priority"]
	var adsPriority int
	if !ok || len(priority[0]) < 1 {
		adsPriority = 10
	}else {
		top, err := strconv.Atoi(priority[0])
		if err == nil {
			adsPriority = top
		}
	}

	// Fetch the unexpired 10 ads of a specific priority sent in request
	ads, err := mdao.FindAdsOfPriority(adsPriority, time.Now())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	rand.Seed(time.Now().Unix())

	// only if no ad is returned from database then
	// add a two random ads from the same priority
	if len(ads) != 0{
		n := rand.Int() % len(ads)
		mainNewsStruct.AdsItem = append(mainNewsStruct.AdsItem, ads[n])
		n2 := rand.Int() % len(ads)
		mainNewsStruct.AdsItem = append(mainNewsStruct.AdsItem, ads[n2])
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(mainNewsStruct); err != nil {
		panic(err)
	}
}

// Function to send 20 ads in order.
// First 5 ads will be of priority 11 and then in order
func GetAdsInOrder(w http.ResponseWriter, r *http.Request) {

	mainNewsStruct := AdsItems{Status:"OK"}

	// Fetch the ads of priority of 11 for top 5 places in ads section
	ads, err := mdao.FindAdsOfPriority(11, time.Now())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	rand.Seed(time.Now().Unix())
	count := 0

	// create a map to check if ad is added
	visitedAd := map[string]bool {}

	// check if ads are returned
	if len(ads) != 0{
		for count < len(ads) { // show priority 11 ads fetched from database ordered by start date
			//n := rand.Int() % len(ads)

			// randomly select an ad of priority 11 and put it in response.
			if(!visitedAd[ads[count].Id.String()]){ // if ad is already not added then add it to avoid repeated ads
				mainNewsStruct.AdsItem = append(mainNewsStruct.AdsItem, ads[count])
				visitedAd [ads[count].Id.String()] = true
			}
			count += 1
		}
	}

	// Now get ads of priority of 12 and above
	moreads, err := mdao.FindAdsAbovePriority(11, time.Now())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	count = 0 // reset counter to traverse through priority > 11 ads
	if len(moreads) != 0{
		for count < len(moreads) {
			//n := rand.Int() % len(moreads)

			// randomly select an ad of priority above 11 and put it in response.
			if(!visitedAd[moreads[count].Id.String()]) { // if ad is already not added then add it
				mainNewsStruct.AdsItem = append(mainNewsStruct.AdsItem, moreads[count])
				visitedAd [moreads[count].Id.String()] = true
			}
			count += 1
		}
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(mainNewsStruct); err != nil {
		panic(err)
	}
}


func PostAds(w http.ResponseWriter, r *http.Request) {

	var ads Campaigns
	if err := json.NewDecoder(r.Body).Decode(&ads); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}
	ads.Id = bson.NewObjectId()
	ads.StartDate = time.Now()// start now

	// put default expiry date of one month
	ads.EndDate = time.Now().AddDate(0, 1, 0) // end after one month


	if err := mdao.InsertAds(ads); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	respondWithAdsJson(w, http.StatusCreated, ads)
}

func DeleteAds(w http.ResponseWriter, r *http.Request) {
	var ad Campaigns

	if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}

	if err := mdao.DeleteAds(ad); err != nil {
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

func respondWithAdsJson(w http.ResponseWriter, status int, ads Campaigns){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(ads); err != nil {
		panic(err)
	}
}

