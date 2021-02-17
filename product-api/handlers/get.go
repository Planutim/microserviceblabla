package handlers

import (
	"context"
	"net/http"

	"github.com/Planutim/microserviceblabla/product-api/data"

	protos "github.com/Planutim/microserviceblabla/currency"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 200: productsResponse

//ListAll handles GET requests returns the all products from the data store
func (p *Products) ListAll(w http.ResponseWriter, req *http.Request) {
	p.l.Println("[DEBUG] get all records")

	w.Header().Add("Content-Type", "application/json")
	prods := data.GetProducts()

	err := data.ToJSON(prods, w)

	if err != nil {
		// http.Error(w, "[ERROR] serializing product", http.StatusInternalServerError)
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from a database
// responses:
// 	200: productResponse
//  404: errorResponse
func (p *Products) ListSingle(w http.ResponseWriter, req *http.Request) {
	id := getProductID(req)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}
	//get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["GBP"]),
	}
	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("[Error]  error getting new rate", err)

		data.ToJSON(&GenericError{Message: err.Error()}, w)
	}
	prod.Price = prod.Price * resp.Rate
	err = data.ToJSON(prod, w)
	if err != nil {
		// we should never expect this
		p.l.Println("[ERROR] serializing product", err)
	}
}
