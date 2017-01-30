package main

import (
	"net/http"
	"github.com/labstack/gommon/log"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./starting-files"))))
}
