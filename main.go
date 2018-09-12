package main

import _ "net/http/pprof"
import (
"log"
"net/http"
"github.com/jaytailor/news-server/handler"
//"github.com/gorilla/context"
	"github.com/gorilla/mux"
)
import "net/http/pprof"

const (
	STATIC_DIR= "/static"
	PORT= "8080"
)

func AttachProfiler(router *mux.Router) {
    router.HandleFunc("/debug/pprof/", pprof.Index)
    router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    router.HandleFunc("/debug/pprof/profile", pprof.Profile)
    router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
  
    // Manually add support for paths linked to by index page at /debug/pprof/
    router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
    router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
    router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
    router.Handle("/debug/pprof/block", pprof.Handler("block"))

}

func main() {
	router := handler.NewRouter() 
	
	// in case we need to hprof again 
	//AttachProfiler(router)

	// To return static images
	InitHandler(router, "static/")

	log.Fatal(http.ListenAndServe(":"+PORT, router))

}


func InitHandler(router *mux.Router, path string)  {
	router.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir(path))))
	}


