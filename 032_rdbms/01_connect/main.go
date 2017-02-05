package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

var db *sql.DB
var err error

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// AWS
	//db, err = sql.Open(
	//	"mysql",
	//	"root:@tcp(Host/IP:3306)/test02?charset=utf8")

	// GCP
	db, err = mysql.DialPassword("CONNECTION INSTANCE NAME", "root", "Password")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
