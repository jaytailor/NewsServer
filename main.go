package main

import (
"log"
"net/http"
"github.com/jaytailor/news-server/handler"
//"github.com/gorilla/context"
	"github.com/gorilla/mux"
)


const (
	STATIC_DIR= "/static"
	PORT= "8080"
)

func main() {
	router := handler.NewRouter()

	// To return static images
	InitHandler(router, "static/")

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}


func InitHandler(router *mux.Router, path string)  {
	router.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir(path))))
	}


