package data

import (
	"encoding/json"
	"io"
	"time"
)

// models a mock table from a database
// extra json declaration to change field names -> needed because fields are capital to be public
type Product struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float32 `json:"price"`
	SKU         string `json:"sku"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

type Products []*Product // list of product pointers

// method called by handler => equiv to marshal
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// returns the actual slice of struct ptrs, not a json string
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "latte",
		Description: "frothy milk coffee",
		Price: 2.45,
		SKU: "abc123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "expresso",
		Description: "short coffee without milk",
		Price: 1.99,
		SKU: "asdf123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}