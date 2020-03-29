package product

import (
	"log"

	"github.com/sylvioneto/store-app/pkg/infra"
)

// Product is the structure that defines store item
type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

// QueryAll returns all products in the database
func QueryAll() []Product {
	db := infra.ConnectToDatabase()
	defer db.Close()

	rows, err := db.Query("select id, name, price, quantity from product")
	if err != nil {
		log.Fatalln(err.Error())
	}

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(product)
		products = append(products, product)
	}
	return products
}

//Save function inserts a product into the table
func (p *Product) Save() {
	db := infra.ConnectToDatabase()
	defer db.Close()
	stmt, err := db.Prepare("insert into product(name, price, quantity) values($1,$2,$3)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(p.Name, p.Price, p.Quantity)
	if err != nil {
		log.Fatalln(err)
	}
}

// Delete a product
func (p *Product) Delete() {
	log.Println(p.ID)
	db := infra.ConnectToDatabase()
	defer db.Close()
	
	stmt, err := db.Prepare("delete from product where id = $1")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(p.ID)
	if err != nil {
		log.Fatalln(err)
	}
}
