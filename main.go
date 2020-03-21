package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sylvioneto/store-app/pkg/items"
)

var webTemplates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	log.Println("Starting webserver...")
	connectToDatabase()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	items := items.GetTestData()
	webTemplates.ExecuteTemplate(w, "index.html", items)
}

func connectToDatabase() {
	connStr := "dbname=storedb host=localhost port=54320 sslmode=disable user=postgres password=mysecretpassword"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
	log.Println("connect to db: success")
}
