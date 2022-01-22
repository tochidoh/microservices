package data

import "fmt"

// ErrProductNotFound raised when product not found in database
var ErrProductNotFound = fmt.Errorf("product not found")

// Product defines structure for api product
// swagger:model
type Product struct {
	// id of the product
	//
	// required: false
	// min: 1
	ID int `json:"id"`

	// name of product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// description of product
	//
	// required: false
	// max length: 10000
	Description string `json:"description`

	// price of product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// SKU for product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

// Products defines a slice of Product pointers
type Products []*Product

// GetProducts returns all products from database
func GetProducts() Products {
	return productList
}

// GetProductById returns a single product which matches the id from database
// if a product is not found, return ErrProductNotFound error
func GetProductById(id int) (*Product, error) {
	index := findIndexByProductId(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[index], nil
}

// UpdateProduct replaces a product in database with given item
// if product with given id does not exist in database return ErrProductNotFound
func UpdateProduct(product Product) error {
	index := findIndexByProductId(product.ID)
	if index == -1 {
		return ErrProductNotFound
	}

	// update product in database
	productList[index] = &product // passed product not *product

	return nil
}

// AddProduct adds a new product to database
func AddProduct(product Product) {
	// get next id in sequence
	maxID := productList[len(productList)-1].ID
	product.ID = maxID + 1
	productList = append(productList, &product)
}

// findIndexByProductId finds the index of produce in database
// returns -1 when no product is found
func findIndexByProductId(id int) int {
	for index, product := range productList {
		if product.ID == id {
			return index
		}
	}

	return -1
}

// DeleteProduct deletes a product from database
// if product with given id does not exist in database return ErrProductNotFound
func DeleteProduct(id int) error {
	index := findIndexByProductId(id)
	if index == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:index], productList[index+1:]...)

	return nil
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
