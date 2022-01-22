package handlers

import (
	"net/http"

	"github.com/tochidoh/microservices/product_api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// update a products details
//
// responses:
// 	201: noContentResponse
// 	404: errorResponse
// 	501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductId(r)

	p.logger.Println("debug: deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.logger.Println("error: deleting record is does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.logger.Println("error: deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
