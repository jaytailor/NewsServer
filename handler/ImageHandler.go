package handler

import (
	"net/http"
	"fmt"
	"io"
	"os"
	"mime/multipart"
	"encoding/json"
)

const (
	STATIC_DIR = "/static/"
)

type res struct {
	msg string
}

func PostImage(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("posting the image now")

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("image")

	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "%v", handler.Header)

	mimeType := handler.Header.Get("Content-Type")

	switch mimeType {
		case "image/jpeg":
			saveFile(w, file, handler)
		case "image/png":
			saveFile(w, file, handler)
		default:
			jsonResponse(w, http.StatusBadRequest, res{msg:"Bad Request"})
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {

	f, err := os.OpenFile("./static/"+handle.Filename, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	defer f.Close()
	_, err = io.Copy(f, file)

	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	message := res{msg:"Success"}
	jsonResponse(w, http.StatusCreated, message)
}

func jsonResponse(w http.ResponseWriter, code int, message res) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(message.msg); err != nil {
		panic(err)
	}
}