package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/sylvioneto/store-app/pkg"
)

var webTemplates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	log.Println("Starting webserver...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	items := items.GetTestData()
	webTemplates.ExecuteTemplate(w, "index.html", items)
}
