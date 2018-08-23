package handler

import (
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetAllNews",
		"GET",
		"/getallnews",
		GetAllNews,
	},
	Route{
		"PostNews",
		"POST",
		"/postnews",
		PostNews,
	},
	Route{
		"DeleteNews",
		"POST",
		"/deletenews",
		DeleteNews,
	},
	Route{
		"PostAds",
		"POST",
		"/postads",
		PostAds,
	},
	Route{
		"GetAllAds",
		"GET",
		"/getallads",
		GetAllAds,
	},
	Route{
		"DeleteAds",
		"POST",
		"/deleteads",
		DeleteAds,
	},
	Route{
		"PostVideos",
		"POST",
		"/postvideos",
		PostVideos,
	},
	Route{
		"GetAllVideos",
		"GET",
		"/getallvideos",
		GetAllVideos,
	},
	Route{
		"PostEditorial",
		"POST",
		"/posteditorial",
		PostEditorial,
	},
	Route{
		"GetAllEditorial",
		"GET",
		"/getalleditorial",
		GetAllEditorial,
	},
	Route{
		"GetImage",
		"GET",
		"/getimage",
		GetImage,
	},

}


