package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sylvioneto/store-app/pkg/routes"
)

func main() {
	log.Println("Starting webserver...")
	routes.SetupRoutes()
	http.ListenAndServe(":8000", nil)
}
