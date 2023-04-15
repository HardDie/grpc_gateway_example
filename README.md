# Example of grpc HTTP gateway

## Install dependency
```bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Build services
```bash
cd appfirst && make build && cd ..
cd appsecond && make build && cd ..
```

## Third party proto files
The files in **gateway/api/proto/google** were copied from https://github.com/googleapis/googleapis.git and are located at **google/api**

The files in **gateway/api/proto/protoc-gen-openapiv2** were copied from https://github.com/grpc-ecosystem/grpc-gateway and are located at **protoc-gen-openapiv2/options**
