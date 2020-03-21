package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sylvioneto/store-app/pkg/product"
)

var webTemplates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	log.Println("Starting webserver...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := product.QueryAll()
	webTemplates.ExecuteTemplate(w, "Index", products)
}
