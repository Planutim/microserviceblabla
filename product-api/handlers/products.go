package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Planutim/microserviceblabla/product-api/data"

	protos "github.com/Planutim/microserviceblabla/product-api/currency"
	"github.com/gorilla/mux"
)

// Products handler for getting and updating products
type Products struct {
	l  *log.Logger
	v  *data.Validation
	cc protos.CurrencyClient
}

// returns new instance of Products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation, cc protos.CurrencyClient) *Products {
	return &Products{l: l, v: v}
}

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(req *http.Request) int {
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	return id
}
