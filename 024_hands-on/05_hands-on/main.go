package main

import (
	"net/http"
	"log"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", hoge)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hoge(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w,nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}