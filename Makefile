generate:
	@go generate ./...
gen_api_doc:
	swag init -d ./cmd/pangud,./internal/sslcert/resource -o ./api --parseDependency