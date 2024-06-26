// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/api.proto

package proto

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

// HomeServiceClient is the client API for HomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeServiceClient interface {
	// Sends a greeting
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type homeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeServiceClient(cc grpc.ClientConnInterface) HomeServiceClient {
	return &homeServiceClient{cc}
}

func (c *homeServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/homeserver.proto.HomeService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeServiceServer is the server API for HomeService service.
// All implementations must embed UnimplementedHomeServiceServer
// for forward compatibility
type HomeServiceServer interface {
	// Sends a greeting
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	mustEmbedUnimplementedHomeServiceServer()
}

// UnimplementedHomeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHomeServiceServer struct {
}

func (UnimplementedHomeServiceServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedHomeServiceServer) mustEmbedUnimplementedHomeServiceServer() {}

// UnsafeHomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeServiceServer will
// result in compilation errors.
type UnsafeHomeServiceServer interface {
	mustEmbedUnimplementedHomeServiceServer()
}

func RegisterHomeServiceServer(s grpc.ServiceRegistrar, srv HomeServiceServer) {
	s.RegisterService(&HomeService_ServiceDesc, srv)
}

func _HomeService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homeserver.proto.HomeService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeService_ServiceDesc is the grpc.ServiceDesc for HomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "homeserver.proto.HomeService",
	HandlerType: (*HomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _HomeService_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}
