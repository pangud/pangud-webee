generate:
	@go generate ./...
gen_api_doc:
	swag init -d ./cmd/pangud,./internal/account/resource -o ./api --parseDependency