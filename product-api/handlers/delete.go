package handlers

import (
	"net/http"
	"strconv"

	"github.com/Planutim/microserviceblabla/product-api/data"

	"github.com/gorilla/mux"
)

//	swagger:route DELETE /products/{id} products deleteProduct
//	Returns result of deletion
//	responses:
//		201: noContentResponse
//		404: errorResponse
//		501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id := vars["id"]
	idInt, _ := strconv.Atoi(id)
	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(idInt)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}
	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
