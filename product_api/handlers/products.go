package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tochidoh/microservices/product_api/data"
)

// KeyProduct is a key used for Product object in context
type KeyProduct struct{}

// Products handler for getting and updating products
type Products struct {
	logger     *log.Logger
	validation *data.Validation
}

// NewProducts returns a new products handler with given logger
func NewProducts(logger *log.Logger, validation *data.Validation) *Products {
	return &Products{logger, validation}
}

// ErrInvalidProductPath is an error message when product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validaiton error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductId returns the product ID from the URL
// panics if cannot convert id into an integer
// this should not happen bc router ensures it's a valid number
func getProductId(r *http.Request) int {
	// parse product id from url
	vars := mux.Vars(r)

	// convert id into integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}
