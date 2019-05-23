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
		"GetAdsWithPriority",
		"GET",
		"/topads",
		GetAdsWithPriority,
	},
	Route{
		"GetAdsInOrder",
		"GET",
		"/ads",
		GetAdsInOrder,
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
		"PostImage",
		"POST",
		"/postimage",
		PostImage,
	},
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
	Route{
		"CreateUser",
		"POST",
		"/createuser",
		CreateUser,
	},
	Route{
		"PostWishMessage",
		"POST",
		"/postwish",
		PostWishMessage,
	},
	Route{
		"GetAllWishes",
		"GET",
		"/getallwishes",
		GetAllWishes,
	},
	Route{
		"PostSurvey",
		"POST",
		"/postsurvey",
		PostSurvey,
	},
	Route{
		"GetSurveyResult",
		"GET",
		"/getsurveyresult",
		GetSurveyResult,
	},
	Route{
		"UpdateSurvey",
		"POST",
		"/updatesurvey",
		UpdateSurvey,
	},
	Route{
		"GetBreakingNews",
		"GET",
		"/breakingnews",
		GetNumBNews,
	},
	Route{
		"PostBreakingNews",
		"POST",
		"/postbreakingnews",
		PostBreakingNews,
	},

}


