package main

import (
	"database/sql"
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
	products := queryAllProducts()
	webTemplates.ExecuteTemplate(w, "index.html", products)
}

func connectToDatabase() *sql.DB {
	connStr := "dbname=storedb host=localhost port=54320 sslmode=disable user=postgres password=mysecretpassword"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("connect to db: success")
	return db
}

func queryAllProducts() []product.Product {
	db := connectToDatabase()
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
