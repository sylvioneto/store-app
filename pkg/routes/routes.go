package routes

import (
	"github.com/sylvioneto/store-app/pkg/product"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", product.Index)
}
