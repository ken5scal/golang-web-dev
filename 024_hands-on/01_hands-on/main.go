package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}