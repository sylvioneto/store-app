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

func queryAll() []Product {
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

func queryByID(ID int) Product {
	db := infra.ConnectToDatabase()
	defer db.Close()

	rows, err := db.Query("select id, name, price, quantity from product where id=$1", ID)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var product Product
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(product)
	}
	return product
}

func (p *Product) save() {
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

func (p *Product) delete() {
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

func (p *Product) update() {
	db := infra.ConnectToDatabase()
	defer db.Close()
	stmt, err := db.Prepare("update product set name=$1, price=$2, quantity=$3 where id=$4")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(p.Name, p.Price, p.Quantity, p.ID)
	if err != nil {
		log.Fatalln(err)
	}
}
