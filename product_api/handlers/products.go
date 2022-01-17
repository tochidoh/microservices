package handlers

import (
	"log"
	"net/http"

	"github.com/tochidoh/microservices/product_api/data"
)

// struct for handler, not actual products
type Products struct {
	l *log.Logger
}

// constructor creates products handler instance => logger is passed from main
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// handler method controls all logic
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// "service" method that gets data, would realistically be a database call
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle get products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw) // same as marshal but writes to response
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("triggering add product")

	// now need to convert json body to normal product struct
	prod := &data.Product{}

	err := prod.FromJSON(r.Body) // request body is a io reader
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("product: %v\n", prod)
}

// encode forms json from struct
// decode forms struct from json
