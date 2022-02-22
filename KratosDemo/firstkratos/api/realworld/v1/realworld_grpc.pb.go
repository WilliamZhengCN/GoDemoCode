// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/realworld/v1/realworld.proto

package v1

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

// RealworldClient is the client API for Realworld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealworldClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type realworldClient struct {
	cc grpc.ClientConnInterface
}

func NewRealworldClient(cc grpc.ClientConnInterface) RealworldClient {
	return &realworldClient{cc}
}

func (c *realworldClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealworldServer is the server API for Realworld service.
// All implementations must embed UnimplementedRealworldServer
// for forward compatibility
type RealworldServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedRealworldServer()
}

// UnimplementedRealworldServer must be embedded to have forward compatible implementations.
type UnimplementedRealworldServer struct {
}

func (UnimplementedRealworldServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedRealworldServer) mustEmbedUnimplementedRealworldServer() {}

// UnsafeRealworldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealworldServer will
// result in compilation errors.
type UnsafeRealworldServer interface {
	mustEmbedUnimplementedRealworldServer()
}

func RegisterRealworldServer(s grpc.ServiceRegistrar, srv RealworldServer) {
	s.RegisterService(&Realworld_ServiceDesc, srv)
}

func _Realworld_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Realworld_ServiceDesc is the grpc.ServiceDesc for Realworld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Realworld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realworld.v1.Realworld",
	HandlerType: (*RealworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Realworld_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/realworld/v1/realworld.proto",
}
