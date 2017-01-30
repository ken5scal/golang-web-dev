package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dog2)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl , err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	tpl.ExecuteTemplate(w, "dog.gohtml",nil)

}

func dog2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}