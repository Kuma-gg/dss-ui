package main

import (
	"net/http"
	"text/template"
)

//User structure
type User struct {
	ID        int
	Name      string
	Firstname string
	Lastname  string
	Email     string
}

//Document structure
type Document struct {
	ID   string
	Name string
	Size int64
}

func indexPage(writer http.ResponseWriter, r *http.Request) {
	type WebData struct {
		Title     string
		Users     []User
		Documents []Document
	}

	//FAKE DATA
	data := WebData{
		Title: "Kuma | Mastery",
		Users: getUsers(),
		Documents: getDocuments(),
	}

	t, err := template.ParseFiles("./views/index.html")

	if err != nil {
		panic(err)
	}

	err = t.Execute(writer, data)
	if err != nil {
		panic(err)
	}

}
