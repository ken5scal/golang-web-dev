package main

import (
	"net/http"
	"strconv"
)

var counter int

func main() {
	http.HandleFunc("/", incrementCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func incrementCookie(res http.ResponseWriter, req *http.Request) {
	counter++

	http.SetCookie(res, &http.Cookie{
		Name: "counter",
		Value: strconv.Itoa(counter),
	})
}