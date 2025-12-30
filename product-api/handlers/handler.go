package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST products")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %v", product)
	data.AddProduct(product)
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "couldnt convert to json", http.StatusInternalServerError)
		return
	}
}

func (p Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "couldnt convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT products")
	product := r.Context().Value(KeyProduct{}).(*data.Product)

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

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := &data.Product{}
		err := product.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
			return
		}

		err = product.Validate()
		if err != nil {
			p.l.Println("Error validating product", err)
			http.Error(w, fmt.Sprintf("error validating product: %v", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
