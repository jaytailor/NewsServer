package main

import (
"log"
"net/http"
"github.com/jaytailor/news-server/handler"
//"github.com/gorilla/context"
)


func main() {
	router := handler.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}



