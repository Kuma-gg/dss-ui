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
	Age       int
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

	data := WebData{
		Title: "Kuma | Mastery",
		Users: []User{
			{ID: 1, Name: "Lolpez", Firstname: "Luis", Lastname: "Lopez", Age: 25, Email: "luis@gmail.com"},
			{ID: 2, Name: "Miguel", Firstname: "Miguel", Lastname: "Lopez", Age: 25, Email: "miguel@gmail.com"},
			{ID: 3, Name: "Jose", Firstname: "Jose", Lastname: "Lopez", Age: 25, Email: "jose@gmail.com"},
		},
		Documents: []Document{
			{ID: "abc123", Name: "abc.txt", Size: 50471},
			{ID: "hdp666", Name: "hdp.js", Size: 74125},
			{ID: "lol420", Name: "lol.jpg", Size: 853857},
		},
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
