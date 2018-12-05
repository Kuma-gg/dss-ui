package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func getUsers() []User {
	var arrayUsers []User
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT id, email,first_name,last_name,name FROM users ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		var email string
		var nameUser string
		err = rows.Scan(&id, &email, &firstName, &lastName, &nameUser)
		if err != nil {
			// handle this error
			panic(err)
		}
		user := User{
			ID: id, Name: nameUser, Firstname: firstName, Lastname: lastName, Email: email,
		}
		arrayUsers = append(arrayUsers, user)
		//fmt.Println(id, firstName)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return arrayUsers
}

func getDocuments() []Document {
	var arrayDocument []Document
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT id, name,size FROM documents ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var size int64
		err = rows.Scan(&id, &name, &size)
		if err != nil {
			// handle this error
			panic(err)
		}
		user := Document{
			ID: id, Name: name, Size: size,
		}
		arrayDocument = append(arrayDocument, user)
		//fmt.Println(id, name)

	}
	db.Close()
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return arrayDocument
}

func saveUser(user User) bool {
	//keys := mux.Vars(request)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := fmt.Sprintf(`  INSERT INTO users ( name,email, first_name, last_name)
						VALUES ( '%s', '%s', '%sn', '%s')`, user.Name, user.Email, user.Firstname, user.Lastname)
	_, err = db.Exec(sqlStatement)
	db.Close()
	if err != nil {
		return false
	}
	return true
}

func deleteUser(user User) bool {
	//keys := mux.Vars(request)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := fmt.Sprintf(` DELETE FROM  users WHERE id = %s`, strconv.Itoa(user.ID))
	_, err = db.Exec(sqlStatement)
	db.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func saveDocument(document Document) bool {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := fmt.Sprintf(`  INSERT INTO documents ( name,size)
						VALUES ( '%s', %s)`, document.Name, strconv.FormatInt(document.Size, 10))
	_, err = db.Exec(sqlStatement)
	db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func deleteDocument(document Document) bool {
	//keys := mux.Vars(request)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := fmt.Sprintf(` DELETE FROM  documents WHERE id = %s`, document.ID)
	_, err = db.Exec(sqlStatement)
	db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func getUserById(id string) Document {
	var user Document
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	query  := "SELECT id, name,size FROM documents where id = '" + id+"'"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var size int64
		err = rows.Scan(&id, &name, &size)
		if err != nil {
			// handle this error
			panic(err)
		}
		user = Document{
			ID: id, Name: name, Size: size,
		}
	}
	db.Close()
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return user
}
