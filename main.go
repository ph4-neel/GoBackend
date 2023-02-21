package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl *template.Template
var db *sql.DB

type studentInfo struct {
	Name   string
	course string
}

func getSql() *sql.DB {

	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/go?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection established")
	return db

}

func init() {

	tmpl = template.Must(template.ParseFiles("crud.html"))
}

func crudHandler(w http.ResponseWriter, r *http.Request) {

	db = getSql()

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	Student := &studentInfo{

		Name:   r.FormValue("name"),
		course: r.FormValue("course"),
	}

	//insert student

	if r.FormValue("submit") == "insert" {

		_, err := db.Exec("insert into register(name,course) values(?,?)", Student.Name, Student.course)

		if err != nil {

			log.Fatal(err)
			tmpl.Execute(w, struct {
				success bool
				message string
			}{success: true, message: err.Error()})
		} else {

			tmpl.Execute(w, struct {
				success bool
				message string
			}{success: true, message: "Record inserted"})

		}

	} else if r.FormValue("submit") == "read" {

		tmpl.Execute(w, struct {
			success bool
			message string
		}{success: true, message: "welcome to user list"})

	}

}

func main() {

	fmt.Println("Application is Started")
	http.HandleFunc("/", crudHandler)
	http.ListenAndServe(":8000", nil)

}
