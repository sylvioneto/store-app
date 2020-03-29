package routes

import (
	"net/http"

	"github.com/sylvioneto/store-app/pkg/product"
)

func SetupRoutes() {
	http.HandleFunc("/", product.Index)
	http.HandleFunc("/product/new", product.New)
	http.HandleFunc("/product/save", product.Save)
	http.HandleFunc("/product/delete", product.Delete)
	http.HandleFunc("/product/edit", product.Edit)
}
