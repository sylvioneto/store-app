package product

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	webTemplates := template.Must(template.ParseGlob("templates/*.html"))
	products := QueryAll()
	webTemplates.ExecuteTemplate(w, "Index", products)
}
