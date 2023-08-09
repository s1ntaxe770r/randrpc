// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/rand/rand.proto

package rand

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

// RandServiceClient is the client API for RandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandServiceClient interface {
	Rand(ctx context.Context, in *RandRequest, opts ...grpc.CallOption) (*RandResponse, error)
}

type randServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRandServiceClient(cc grpc.ClientConnInterface) RandServiceClient {
	return &randServiceClient{cc}
}

func (c *randServiceClient) Rand(ctx context.Context, in *RandRequest, opts ...grpc.CallOption) (*RandResponse, error) {
	out := new(RandResponse)
	err := c.cc.Invoke(ctx, "/RandService/Rand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandServiceServer is the server API for RandService service.
// All implementations must embed UnimplementedRandServiceServer
// for forward compatibility
type RandServiceServer interface {
	Rand(context.Context, *RandRequest) (*RandResponse, error)
	mustEmbedUnimplementedRandServiceServer()
}

// UnimplementedRandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRandServiceServer struct {
}

func (UnimplementedRandServiceServer) Rand(context.Context, *RandRequest) (*RandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rand not implemented")
}
func (UnimplementedRandServiceServer) mustEmbedUnimplementedRandServiceServer() {}

// UnsafeRandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandServiceServer will
// result in compilation errors.
type UnsafeRandServiceServer interface {
	mustEmbedUnimplementedRandServiceServer()
}

func RegisterRandServiceServer(s grpc.ServiceRegistrar, srv RandServiceServer) {
	s.RegisterService(&RandService_ServiceDesc, srv)
}

func _RandService_Rand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandServiceServer).Rand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RandService/Rand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandServiceServer).Rand(ctx, req.(*RandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RandService_ServiceDesc is the grpc.ServiceDesc for RandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RandService",
	HandlerType: (*RandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Rand",
			Handler:    _RandService_Rand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rand/rand.proto",
}
