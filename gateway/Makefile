.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: proto
proto: ## Generate proto
	go mod vendor
	protoc -I./pkg/gateway_service \
		-I./api/proto \
		-I./vendor/github.com/HardDie/grpc_gateway_example/appfirst/pkg/appfirst_service \
		-I./vendor/github.com/HardDie/grpc_gateway_example/appsecond/pkg/appsecond_service \
		--go_out ./pkg/gateway_service \
		--go_opt=paths=source_relative \
		--go-grpc_out ./pkg/gateway_service \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out ./pkg/gateway_service \
		--grpc-gateway_opt=paths=source_relative \
		--openapiv2_out ./api \
		--openapiv2_opt allow_merge=true,merge_file_name=api,output_format=yaml \
		./pkg/gateway_service/*.proto

.PHONY: hello
hello: ## Example call Hello methods
	curl localhost:9003/api/v1/first
	curl localhost:9003/api/v1/second

.PHONY: build
build: ## Build binary file
	make proto
	go build -o gateway cmd/main.go
