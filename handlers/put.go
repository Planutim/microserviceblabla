package handlers

import (
	"example/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
// 201: noContentResponse
// 404: errorResponse
// 422: errorValidation

// Update handles PUT requests to update products
func (p *Products) UpdateProducts(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	// fetch the product from the ctx
	prod := req.Context().Value(KeyProduct{}).(*data.Product)

	prod.ID = id
	p.l.Printf("Prod: %#v", prod.ID)
	errN := data.UpdateProduct(*prod)
	if errN == data.ErrProductNotFound {
		p.l.Println(w, "[ERROR] product not found", errN)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, w)
	}

	// write the no content success header
	w.WriteHeader(http.StatusNoContent)

}
