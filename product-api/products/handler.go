package handlers

import (
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST products")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %v", product)
	data.AddProduct(product)
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "couldnt convert to json", http.StatusInternalServerError)
		return
	}
}

func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "couldnt convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT products")
	product := &data.Product{}

	err = product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
		return
	}

	if err == data.Err1 {
		http.Error(w, "product not found", http.StatusBadRequest)
		return
	}
	data.UpdateProduct(id, product)
}
