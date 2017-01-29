package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

func init() {
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/", root)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "dog")
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "root")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("03_hands-on.html")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "03_hands-on.html", "My Name")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
