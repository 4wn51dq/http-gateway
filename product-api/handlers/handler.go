// Package Classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 0.0.1
//
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// swagger:meta
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

// swagger:route POST /products products addProduct
//
// Adds a new product.
//
// responses:
//
//	201: productResponse
//	400: errorResponse
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

// swagger:route GET /products products listProducts
//
// Returns the list of products.
//
// responses:
//
//	200: productsResponse
//	500: errorResponse
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "couldnt convert to json", http.StatusInternalServerError)
		return
	}
}

// swagger:route PUT /products/{id} products updateProduct
//
// Updates an existing product.
//
// responses:
//
//	200: productResponse
//	400: errorResponse
//	404: errorResponse
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

// swagger:response errorResponse
type ErrorResponse struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:response productsResponse
type ProductsResponse struct {
	// in: body
	Body []*data.Product
}

// swagger:response productResponse
type ProductResponseWrapper struct {
	// in: body
	Body data.Product
}

// swagger:parameters updateProduct
type UpdateProductParams struct {
	// in: path
	// required: true
	ID int `json:"id"`

	// in: body
	// required: true
	Body *data.Product
}
