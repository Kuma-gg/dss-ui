package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func userSave(writer http.ResponseWriter, req *http.Request) {
	//Decode JSON
	userJSON, err := json.Marshal(User{
		ID:        1,
		Name:      req.FormValue("name"),
		Firstname: req.FormValue("first_name"),
		Lastname:  req.FormValue("last_name"),
		Email:     req.FormValue("email"),
	})
	if err != nil {
		panic(err)
	}
	log.Print(userJSON)

	//Decode JSON
	var userNormal User
	errDecoding := json.Unmarshal(userJSON, &userNormal)
	if errDecoding != nil {
		panic(errDecoding)
	}
	log.Print(userNormal.ID)
	log.Print(userNormal.Name)
	log.Print(userNormal.Firstname)
	log.Print(userNormal.Lastname)
	log.Print(userNormal.Email)
	http.Redirect(writer, req, "/", http.StatusMovedPermanently)
}

func userDelete(writer http.ResponseWriter, r *http.Request) {

}
