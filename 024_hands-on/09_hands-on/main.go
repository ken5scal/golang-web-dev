package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", dogs)
	http.Handle("/public/", http.StripPrefix("/public", fs))
	log.Fatalln(http.ListenAndServe(":8080",nil))
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
