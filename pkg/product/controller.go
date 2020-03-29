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
	products := queryAll()
	webTemplates.ExecuteTemplate(w, "Index", products)
}

// New redirects to the new page
func New(w http.ResponseWriter, r *http.Request) {
	webTemplates.ExecuteTemplate(w, "New", nil)
}

// Save receives the form data and calls save method
func Save(w http.ResponseWriter, r *http.Request) {
	log.Println("Save method")
	if r.Method == "POST" {
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Fatalln(err.Error())
		}
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if r.FormValue("id") == "" {
			p := Product{Name: r.FormValue("name"), Price: price, Quantity: quantity}
			p.save()

		} else {
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				log.Fatalln(err.Error())
			}
			p := Product{ID: id, Name: r.FormValue("name"), Price: price, Quantity: quantity}
			p.update()
		}
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
	product.delete()
	http.Redirect(w, r, "/", 301)
}

// Edit handles edit requests
func Edit(w http.ResponseWriter, r *http.Request) {
	log.Println("Edit handler")
	productID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	product := queryByID(productID)
	webTemplates.ExecuteTemplate(w, "Edit", product)
}
