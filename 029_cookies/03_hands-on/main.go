package main

import "net/http"

var counter int

func main() {
	http.HandleFunc("/", incrementCookie)
	http.ListenAndServe(":8080", nil)
}

func incrementCookie(res http.ResponseWriter, req *http.Request) {

}