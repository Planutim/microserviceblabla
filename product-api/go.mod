module example

go 1.15

replace github.com/currency/protos => /home/leviathan/Documents/currency/protos

require (
	github.com/currency/protos v0.0.0-00010101000000-000000000000
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.10
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0
	google.golang.org/grpc v1.35.0 // indirect
)
