package handlers

import (
	"context"
	"example/data"
	"net/http"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (p *Products) MiddleWareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, req.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		//validate the product after retrieving from JSON
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			// return the validation messages as an array
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&errs, w)
		}
		// add the product to the context
		ctx := context.WithValue(req.Context(), KeyProduct{}, prod)
		req1 := req.WithContext(ctx)

		// call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, req1)
	})
}
