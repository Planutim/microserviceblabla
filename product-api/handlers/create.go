package handlers

import (
	"net/http"

	"github.com/Planutim/microserviceblabla/product-api/data"
)

// swagger:route POST / products productsRequest
// Returns result of creating
// responses:
//	201: noContentResponse
// 422: errorValidation
// 501: errorResponse
func (p *Products) AddProduct(w http.ResponseWriter, req *http.Request) {
	p.l.Println("Handls POST Product")

	prod := req.Context().Value(KeyProduct{}).(*data.Product)
	// prod := &data.Product{}
	// err := prod.FromJSON(req.Body)
	// if err != nil {
	// 	http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	// }

	p.l.Println("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
