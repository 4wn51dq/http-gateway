package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       uint   `json:"price" validate:"gt=0"`
	SKU         string `json:"sku" validate:"required,sku"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

type Products []*Product

func (pl *Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(pl)
}

func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func GetProducts() Products {
	return products
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	products = append(products, p)
}

func UpdateProduct(id int, p *Product) error {
	product, position, err := findProduct(id)
	if err != nil {
		return err
	}
	product.ID = id
	p.ID = id
	products[position] = p
	return nil
}

var Err1 error = fmt.Errorf("couldnt find product with the given id")

func findProduct(id int) (*Product, int, error) {
	for i, p := range products {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, Err1
}

func getNextID() int {
	lp := products[len(products)-1]
	return lp.ID + 1
}

////////////////////////////////////////////

var products = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "frothed milk in espresso",
		Price:       43,
		SKU:         "c0ff33",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Mocha",
		Description: "frothed milk in espresso with cocoa",
		Price:       48,
		SKU:         "c0c04c0ff33",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
