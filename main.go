package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	router := mux.NewRouter()
	/*router.HandleFunc("/user/save", saveUser).Methods("POST")
	router.HandleFunc("/index", userForm).Methods("GET")*/
	router.HandleFunc("/", indexPage).Methods("GET")

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
