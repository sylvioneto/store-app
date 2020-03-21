package product

import (
	"log"
	"github.com/sylvioneto/store-app/pkg/infra"
)

// Product is the structure that defines store item
type Product struct {
	ID int
	Name string
	Price float64
	Quantity int
}

// GetTestData returns mock data
func GetTestData() []Product {
	products := []Product{
		{1, "Rice", 11.99, 5},
		{2, "Water", 0.99, 20},
		{3, "Meat", 29.11, 1},
	}
	return products
}

// QueryAll returns all products in the database
func QueryAll() []Product {
	db := infra.ConnectToDatabase()
	defer db.Close()

	rows, err := db.Query("select id, name, price, quantity from items")
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
