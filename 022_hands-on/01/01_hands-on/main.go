package main

import (
	"net/http"
	"fmt"
)

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "dog")
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "root")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "me")
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/", root)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
