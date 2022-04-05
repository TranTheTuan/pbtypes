// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/auth/auth_verify.proto

package auth

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

// AuthVerifyServiceClient is the client API for AuthVerifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthVerifyServiceClient interface {
	VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
}

type authVerifyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthVerifyServiceClient(cc grpc.ClientConnInterface) AuthVerifyServiceClient {
	return &authVerifyServiceClient{cc}
}

func (c *authVerifyServiceClient) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	out := new(VerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.AuthVerifyService/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthVerifyServiceServer is the server API for AuthVerifyService service.
// All implementations must embed UnimplementedAuthVerifyServiceServer
// for forward compatibility
type AuthVerifyServiceServer interface {
	VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error)
	mustEmbedUnimplementedAuthVerifyServiceServer()
}

// UnimplementedAuthVerifyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthVerifyServiceServer struct {
}

func (UnimplementedAuthVerifyServiceServer) VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedAuthVerifyServiceServer) mustEmbedUnimplementedAuthVerifyServiceServer() {}

// UnsafeAuthVerifyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthVerifyServiceServer will
// result in compilation errors.
type UnsafeAuthVerifyServiceServer interface {
	mustEmbedUnimplementedAuthVerifyServiceServer()
}

func RegisterAuthVerifyServiceServer(s grpc.ServiceRegistrar, srv AuthVerifyServiceServer) {
	s.RegisterService(&AuthVerifyService_ServiceDesc, srv)
}

func _AuthVerifyService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthVerifyServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.AuthVerifyService/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthVerifyServiceServer).VerifyToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthVerifyService_ServiceDesc is the grpc.ServiceDesc for AuthVerifyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthVerifyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.AuthVerifyService",
	HandlerType: (*AuthVerifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyToken",
			Handler:    _AuthVerifyService_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth/auth_verify.proto",
}