package handler

import (
	"net/http"
	"fmt"
)

const (
	STATIC_DIR = "/static/"
)

func GetImage(w http.ResponseWriter, r *http.Request) {

	// Server CSS, JS & Images Statically.
	http.Handle(STATIC_DIR, http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("image.png"))))

	fmt.Println("getting the image now")
	
}