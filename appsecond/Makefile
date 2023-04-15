.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: proto
proto: ## Generate proto
	protoc -I./pkg/appsecond_service \
		--go_out ./pkg/appsecond_service \
		--go_opt=paths=source_relative \
		--go-grpc_out ./pkg/appsecond_service \
		--go-grpc_opt=paths=source_relative \
		./pkg/appsecond_service/*.proto

.PHONY: hello
hello: ## Example call Hello method without reflection
	grpcurl -plaintext \
		-proto pkg/appsecond_service/appsecond-service.proto \
		-d '{}' \
		localhost:9002 \
		appsecond_service.AppSecondService.Hello

.PHONY: build
build: ## Build binary file
	make proto
	go build -o appsecond cmd/main.go
