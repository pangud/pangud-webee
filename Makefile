wire:
	@go generate ./...
gen_api_doc:
	# swag init -d ./cmd/pangud,./internal/sslcert/resource,./internal/core/resource -o ./api --parseDependency
	swag init -d ./cmd/pangud,./internal/sslcert/resource,./internal/core/biz,./internal/core/resource -o ./api
gen_core_data:
	go run internal/core/gen/gorm_gen.go
generate: wire gen_api_doc gen_core_data
	@echo "generate done"