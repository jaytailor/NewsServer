package handler

import (
	"encoding/json"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func Login(w http.ResponseWriter, r *http.Request) {
	mdao := DbDAO{Server:"localhost", Database:LOGIN}
	mdao.Connect()
	//defer session.Close()

	// Create a struct to respond back
	response := Success{Authenticated: "TRUE"}

	// Close the request body in the end
	defer r.Body.Close()

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

		if(err != nil){
			fmt.Println("Error fetching user details")
			panic(err)
			return
		}

		for _, element := range logindetails{
			//editorialStruct.ArticleList = append(editorialStruct.ArticleList, element)

			if(element.User != username[0]) {
				response.Authenticated = "FALSE";
				response.Message = "Username is wrong";
				SendResponse(w, http.StatusOK, response);
				return
			}

			if(element.Password != password[0]){
				response.Authenticated = "FALSE";
				response.Message = "Username or Password is wrong";
				SendResponse(w, http.StatusOK, response);
				return
			}

			response.Authenticated = "TRIE";
			response.Message = "User Authenticated";
			SendResponse(w, http.StatusOK, response);
			return
		}
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	mdao := DbDAO{Server:"localhost", Database:LOGIN}
	mdao.Connect()
	//defer session.Close()

	// Create a struct to respond back
	response := Success{Authenticated: "TRUE", Message:"User Created Successfully"}

	// Close the request body in the end
	defer r.Body.Close()

	// Load new user data from the post data
	var userdata Logins
	if err := json.NewDecoder(r.Body).Decode(&userdata); err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		panic(err)
		return
	}

	newuser := Logins{};

	newuser.Id = bson.NewObjectId()
	newuser.User = userdata.User;
	newuser.Password = userdata.Password;
	newuser.Role = userdata.Role;

	fmt.Println("Creating user now %S %S %S", newuser.User, newuser.Password, newuser.Role);

	if err := mdao.CreateLogins(newuser); err != nil {
		//respondWithError(w, http.StatusInternalServerError, err.Error())
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func SendResponse(w http.ResponseWriter, status int, response Success){

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}