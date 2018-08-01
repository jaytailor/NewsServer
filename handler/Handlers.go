package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
	//"io/ioutil"
	//"io"
	//"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strconv"
)


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, %q", html.EscapeString(r.URL.Path))
}

func GetAllNews(w http.ResponseWriter, r *http.Request) {
	mdao := NewsDAO{Server:"localhost", Database:"news"}
	mdao.Connect()

	keys, ok := r.URL.Query()["list"]
	var number_of_news int
	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		number_of_news = 10
	}else {
		items, err := strconv.Atoi(keys[0])
		if err == nil {
			fmt.Printf("Number of news.. ")
			fmt.Println(items)
			number_of_news = items
		}
	}

	samachar, err := mdao.FindNumOfNews(number_of_news)
	if err != nil {
		//respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(samachar); err != nil {
		panic(err)
	}
}

func PostNews(w http.ResponseWriter, r *http.Request) {
	mdao := NewsDAO{Server:"localhost", Database:"news"}
	mdao.Connect()
	//defer session.Close()

	defer r.Body.Close()
	var samachar NewsModel
	if err := json.NewDecoder(r.Body).Decode(&samachar); err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}
	samachar.Id = bson.NewObjectId()
	samachar.PublishedAt = time.Now()

	if err := mdao.Insert(samachar); err != nil {
		//respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	respondWithJson(w, http.StatusCreated, samachar)
}

func FindSpecificNews(w http.ResponseWriter, r *http.Request) {
	mdao := NewsDAO{Server:"localhost", Database:"news"}
	mdao.Connect()

	params := mux.Vars(r)
	var samachar NewsModel
	samachar, err := mdao.FindById(params["id"])
	if err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		panic(err)
		return
	}
	respondWithJson(w, http.StatusOK, samachar)
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {
	mdao := NewsDAO{Server:"localhost", Database:"news"}
	mdao.Connect()

	defer r.Body.Close()
	var samachar NewsModel
	if err := json.NewDecoder(r.Body).Decode(&samachar); err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}
	if err := mdao.Update(samachar); err != nil {
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

func DeleteNews(w http.ResponseWriter, r *http.Request) {
	mdao := NewsDAO{Server:"localhost", Database:"news"}
	mdao.Connect()

	defer r.Body.Close()
	var samachar NewsModel
	if err := json.NewDecoder(r.Body).Decode(&samachar); err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}
	if err := mdao.Delete(samachar); err != nil {
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

func respondWithJson(w http.ResponseWriter, status int, samachar NewsModel){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(samachar); err != nil {
		panic(err)
	}
}


func ErrorPage(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "TodoShow: ", todoId)
}