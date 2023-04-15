.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: proto
proto: ## Generate proto
	protoc -I./pkg/appfirst_service \
		--go_out ./pkg/appfirst_service \
		--go_opt=paths=source_relative \
		--go-grpc_out ./pkg/appfirst_service \
		--go-grpc_opt=paths=source_relative \
		./pkg/appfirst_service/*.proto

.PHONY: hello
hello: ## Example call Hello method without reflection
	grpcurl -plaintext \
		-proto pkg/appfirst_service/appfirst-service.proto \
		-d '{}' \
		localhost:9001 \
		appfirst_service.AppFirstService.Hello

.PHONY: build
build: ## Build binary file
	make proto
	go build -o appfirst cmd/main.go
