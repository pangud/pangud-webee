wire:
	@go generate ./...
	@echo "wire gen done"
gen_api_doc:
	@swag init -d ./cmd/pangud,./pkg/types,./pkg/errors,./internal/sslcert/resource,\
	./internal/core/biz,./internal/core/resource -o ./api
	@echo "generate api doc done"
gen_core_data:
	@go run internal/core/gen/gorm_gen.go
	@echo "generate core data done"
generate: wire gen_api_doc gen_core_data
	@echo "generate done"
run:
	@go run ./cmd/pangud/... -config configs/config.yaml	