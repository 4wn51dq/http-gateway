package main

import (
	"product-api/client/client"
	"product-api/client/client/products"
	"testing"
)

func TestClient(t *testing.T) {
	c := client.Default
	params := products.NewListProductsParams()
	c.Products.ListProducts(params)
}
