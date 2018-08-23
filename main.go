package main

import (
"log"
"net/http"
"github.com/jaytailor/news-server/handler"
//"github.com/gorilla/context"
)


const (
	STATIC_DIR = "/static/"
	PORT       = "8080"
)

func main() {
	router := handler.NewRouter()
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}



