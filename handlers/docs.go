// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "example/data"

//
// NOTE: types defined here are purely for documentation purposes
// these types are not used by any of the handlers

//Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in:body
	Body data.ValidationError
}

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in:body
	Body []data.Product
}

//Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Get product of a system
	// in:body
	Body data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// A product to add
// swagger:parameters createProduct updateProduct
type productParamsWrapper struct {
	// A product to add from request
	// in:body
	Body data.Product
}

// swagger:parameters updateProduct deleteProduct listSingleProduct
type productsIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in:path
	// required: true
	ID int `json:"id"`
}
