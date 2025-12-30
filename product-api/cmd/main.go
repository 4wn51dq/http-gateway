package main

import (
	"log"
	"net/http"
	"os"
	handlers "product-api/products"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	prodHandler := handlers.NewProducts(l)
	serveMux := http.NewServeMux()
	serveMux.Handle("/coffee", prodHandler)

	http.ListenAndServe(":8000", prodHandler)
}
