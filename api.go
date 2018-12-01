package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)
type WebData struct {
	Table string
}

type UerData struct {
	id string
	name string
}

func userForm(writer http.ResponseWriter, r *http.Request) {

	//--------------

	//--------------
	tmpl := template.Must(template.ParseFiles("forms.html"))
	tableUsers := getUsers()
	wd := WebData{
		Table: tableUsers,
	}

	if r.Method != http.MethodPost {

		tmpl.Execute(writer, &wd)
		return
	}
	//
	tmpl.Execute(writer, struct{ Success bool }{true})

}

func getUsers()  string{
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//var jjjjjj []string

	rows, err := db.Query("SELECT id, email FROM users ")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	tableHtml :=""
	//tableHtml :=" <table style='width:100%'><tr><th>id</th> <th>name</th> </tr> "
	for rows.Next() {
		//tableHtml += "<tr>"

		var id int
		var firstName string
		err = rows.Scan(&id, &firstName)
		if err != nil {
			// handle this error
			panic(err)
		}
		//jjjjjj = append(jjjjjj,string(id))
		tableHtml += " id : "+strconv.Itoa(id)+ "  email : "+firstName+" \n"
		//tableHtml += fmt.Sprintf("<th> %s </th> <th> %s </th>",id,firstName)
		fmt.Println(id, firstName)
		//tableHtml += "</tr>"

	}
	//tableHtml +=" </table>"
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return tableHtml
}
func saveUser(writer http.ResponseWriter, request *http.Request) {
	//keys := mux.Vars(request)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	email := request.FormValue("email")

	sqlStatement := fmt.Sprintf(`  INSERT INTO users (age, name,email, first_name, last_name)
						VALUES (30, 'Jonathan', '%s', 'Jonathan', 'Calhoun')`,email)
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(writer).Encode("se guardo correctamente")
}
