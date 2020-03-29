package product

import (
	"log"
	"net/http"
	"strconv"
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

// Insert receives the form data and calls save method
func Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("Insert method")
	if r.Method == "POST" {
		// convert values
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil{
			log.Fatalln(err.Error())
		}
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil{
			log.Fatalln(err.Error())
		}

		p := Product{Name: r.FormValue("name"), Price: price, Quantity: quantity}
		p.Save()
		log.Println(p)
	}
	http.Redirect(w, r, "/", 301)
}

// Delete receives the id and redirects to database delete function
func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete method")
	productID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	product := Product{ID: productID}
	product.Delete()
	http.Redirect(w, r, "/", 301)
}
