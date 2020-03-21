package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sylvioneto/store-app/pkg/infra"
	"github.com/sylvioneto/store-app/pkg/product"
)

var webTemplates = template.Must(template.ParseGlob("templates/*.html"))
var count int 

func main() {
	log.Println("Starting webserver...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	count++
	//products := queryAllProducts()
	webTemplates.ExecuteTemplate(w, "Index", nil)
	log.Println(count)
}

func queryAllProducts() []product.Product {
	db := infra.ConnectToDatabase()
	defer db.Close()

	rows, err := db.Query("select id, name, price, quantity from items")
	if err != nil {
		log.Fatalln(err.Error())
	}

	var products []product.Product
	for rows.Next() {
		var product product.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(product)
		products = append(products, product)
	}
	return products
}
