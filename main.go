package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var templ *template.Template

type studentInfo struct {
	Name   string
	course string
}

func init() {

	templ = template.Must(template.ParseFiles("crud.html"))
}

func crudHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		templ.Execute(w, nil)
		return
	}

	Student := &studentInfo{

		Name:   r.FormValue("name"),
		course: r.FormValue("course"),
	}

	fmt.Println(Student)
}

func main() {

	fmt.Println("Application is Started")
	http.HandleFunc("/", crudHandler)
	http.ListenAndServe(":8000", nil)

}
