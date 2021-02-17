.DEFAULT_GOAL=swagger
check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger_old:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

codegen:
	swagger generate client -f ./swagger.yaml -A product-api