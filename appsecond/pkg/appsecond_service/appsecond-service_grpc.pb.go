// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: appsecond-service.proto

package appsecond_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AppSecondService_Hello_FullMethodName = "/appsecond_service.AppSecondService/Hello"
)

// AppSecondServiceClient is the client API for AppSecondService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppSecondServiceClient interface {
	// Get hello message
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type appSecondServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppSecondServiceClient(cc grpc.ClientConnInterface) AppSecondServiceClient {
	return &appSecondServiceClient{cc}
}

func (c *appSecondServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, AppSecondService_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppSecondServiceServer is the server API for AppSecondService service.
// All implementations must embed UnimplementedAppSecondServiceServer
// for forward compatibility
type AppSecondServiceServer interface {
	// Get hello message
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedAppSecondServiceServer()
}

// UnimplementedAppSecondServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppSecondServiceServer struct {
}

func (UnimplementedAppSecondServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedAppSecondServiceServer) mustEmbedUnimplementedAppSecondServiceServer() {}

// UnsafeAppSecondServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppSecondServiceServer will
// result in compilation errors.
type UnsafeAppSecondServiceServer interface {
	mustEmbedUnimplementedAppSecondServiceServer()
}

func RegisterAppSecondServiceServer(s grpc.ServiceRegistrar, srv AppSecondServiceServer) {
	s.RegisterService(&AppSecondService_ServiceDesc, srv)
}

func _AppSecondService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppSecondServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppSecondService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppSecondServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppSecondService_ServiceDesc is the grpc.ServiceDesc for AppSecondService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppSecondService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appsecond_service.AppSecondService",
	HandlerType: (*AppSecondServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _AppSecondService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appsecond-service.proto",
}
