package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/user/save", saveUser).Methods("POST")
	router.HandleFunc("/document/save", saveDocument).Methods("POST")
	router.HandleFunc("/index", userForm).Methods("GET")
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) )
	log.Fatal(http.ListenAndServe(":9001", router))
}

func saveDocument(writer http.ResponseWriter, request *http.Request) {


}

