package product

import (
	"net/http"
	"text/template"
)

var webTemplates = template.Must(template.ParseGlob("templates/*.html"))

// Index query products and send to the template
func Index(w http.ResponseWriter, r *http.Request) {
	products := QueryAll()
	webTemplates.ExecuteTemplate(w, "Index", products)
}

// New redirects to the new page
func New(w http.ResponseWriter, r *http.Request) {
	webTemplates.ExecuteTemplate(w, "New", nil)
}
