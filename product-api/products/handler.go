package handlers

import (
	"log"
	"net/http"
	"product-api/data"
	"regexp"
	"strconv"
)

type Product struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w, r)
		return
	case http.MethodPost:
		p.addProducts(w, r)
	case http.MethodPut:
		// expect the id in the uri
		regexp := regexp.MustCompile(`/[0-9]+`) // extracting params out of URI!!
		group := regexp.FindAllStringSubmatch(r.URL.Path, -1)
		if len(group) != 1 {
			http.Error(w, "Invalid uri: more than 1 id", http.StatusBadRequest)
			return
		}
		if len(group[0]) != 1 {
			http.Error(w, "invalid uri: more than 1 capture group", http.StatusBadRequest)
		}
		IDString := group[0][1]
		id, err := strconv.Atoi(IDString)
		if err != nil {
			http.Error(w, "id string could not be converted to int", http.StatusBadRequest)
		}
		p.l.Printf("got the id: %v", id)

		p.updateProduct(id, w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Product) addProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST products")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %v", product)
	data.AddProduct(product)
}

func (p *Product) getProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "couldnt convert to json", http.StatusInternalServerError)
		return
	}
}

func (p *Product) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT products")
	product := &data.Product{}

	err := product.FromJSON(r.Body)
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
