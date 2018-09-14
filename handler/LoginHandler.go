package handler

import (
	"encoding/json"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// Create a struct to respond back
	response := Success{Authenticated: "FALSE"}

	username, ok := r.URL.Query()["username"]
	password, ok := r.URL.Query()["password"]

	if !ok || len(username[0]) < 1 || len(password[0]) < 1 {
		fmt.Println("Username or password is missing")
		response.Authenticated = "INVALID";
		response.Message = "No username password were sent in the request";
		SendResponse(w, http.StatusNotImplemented, response)
		return
	}else {
		logindetails, err := mdao.FindUsers();

		if err != nil{
			fmt.Println("Error fetching user details")
			panic(err)
			return
		}

		if len(logindetails) == 0{
			response.Authenticated = "INVALID";
			response.Message = "No users exist in the system";
			SendResponse(w, http.StatusNotImplemented, response)
			return
		}

		for _, element := range logindetails{

			if element.User != username[0] {
				response.Authenticated = "FALSE";
				response.Message = "Username is wrong";
				SendResponse(w, http.StatusForbidden, response);
				return
			}

			if element.Password != password[0]{
				response.Authenticated = "FALSE";
				response.Message = "Username or Password is wrong";
				SendResponse(w, http.StatusForbidden, response);
				return
			}

			response.Authenticated = "TRUE";
			response.Message = "User Authenticated";
			SendResponse(w, http.StatusOK, response);
		}
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Load new user data from the post data
	var userdata Logins

	// Create a struct to respond back
	response := Success{Authenticated: "TRUE", Message:"User Created Successfully"}

	if err := json.NewDecoder(r.Body).Decode(&userdata); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}

	userdata.Id = bson.NewObjectId()

	fmt.Println("Creating user now ", userdata.User, userdata.Password, userdata.Role);

	if err := mdao.CreateLogins(userdata); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	SendResponse(w, http.StatusOK, response)
}

func SendResponse(w http.ResponseWriter, status int, response Success){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
