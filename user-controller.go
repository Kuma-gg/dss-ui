package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func userSave(writer http.ResponseWriter, req *http.Request) {
	//Encode JSON
	user := User{
		ID:        1,
		Name:      req.FormValue("name"),
		Firstname: req.FormValue("first_name"),
		Lastname:  req.FormValue("last_name"),
		Email:     req.FormValue("email"),
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	saved := saveUser(user)
	if !saved {
		log.Print("Error : Create user!! ")
	}
	//Decode JSON
	var userNormal User
	errDecoding := json.Unmarshal(userJSON, &userNormal)
	if errDecoding != nil {
		panic(errDecoding)
	}

	http.Redirect(writer, req, "/#users", http.StatusMovedPermanently)
}

func userDelete(writer http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	user := User{
		ID: userID,
	}

	//Del user by id
	deleted := deleteUser(user)
	if !deleted {
		log.Print("ERROR : delete User")
	}
	//Encode JSON
	userJSON, err := json.Marshal(user)
	log.Print(userJSON)
	//Decode JSON
	var userNormal User
	errDecoding := json.Unmarshal(userJSON, &userNormal)
	if errDecoding != nil {
		panic(errDecoding)
	}
	log.Print(userNormal.ID)
	http.Redirect(writer, req, "/", http.StatusMovedPermanently)
}

func notifyMail(writer http.ResponseWriter, req *http.Request) {
	//obtener todos los mails en, los leemos del request en el post
	var mails []Mail
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &mails)
	//obtenemos los parseamos y los convertimos a []bytes
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(mails)
	//lo ponemos en la cola

	//sendEmailChannel(reqBodyBytes.Bytes()) // this is the []byte

	sendEmailChannel(reqBodyBytes.Bytes()) // this is the []byte
	http.Redirect(writer, req, "/", http.StatusMovedPermanently)

}
