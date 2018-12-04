package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//ContactDetails struct
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}
var rabbitServer string
var storageRequestQueue string
var storageResponseQueue string
// sql Parameters
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "go_proyect_ui"
)

func main() {
	//RabbitMq server
	rabbitServer = "amqp://guest:guest@localhost:5672/"
	storageRequestQueue = "storageRequestQueue"
	storageResponseQueue = "storageResponseQueue"
	//
	go receiverFileMessage()
	router := mux.NewRouter()
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/user", userSave).Methods("POST")
	router.HandleFunc("/user/delete", userDelete).Methods("POST")
	router.HandleFunc("/document", documentSave).Methods("POST")
	router.HandleFunc("/document/delete", documentDelete).Methods("POST")

	fs := http.FileServer(http.Dir("./public"))
	router.PathPrefix("/js/").Handler(fs)
	router.PathPrefix("/css/").Handler(fs)
	router.PathPrefix("/img/").Handler(fs)
	router.PathPrefix("/fonts/").Handler(fs)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
