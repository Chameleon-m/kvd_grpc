// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: author.proto

package author

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
	Author_GetByBook_FullMethodName = "/kvado.Author/getByBook"
)

// AuthorClient is the client API for Author service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorClient interface {
	GetByBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*AuthorListResponse, error)
}

type authorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorClient(cc grpc.ClientConnInterface) AuthorClient {
	return &authorClient{cc}
}

func (c *authorClient) GetByBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*AuthorListResponse, error) {
	out := new(AuthorListResponse)
	err := c.cc.Invoke(ctx, Author_GetByBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServer is the server API for Author service.
// All implementations must embed UnimplementedAuthorServer
// for forward compatibility
type AuthorServer interface {
	GetByBook(context.Context, *BookRequest) (*AuthorListResponse, error)
	mustEmbedUnimplementedAuthorServer()
}

// UnimplementedAuthorServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorServer struct {
}

func (UnimplementedAuthorServer) GetByBook(context.Context, *BookRequest) (*AuthorListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByBook not implemented")
}
func (UnimplementedAuthorServer) mustEmbedUnimplementedAuthorServer() {}

// UnsafeAuthorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServer will
// result in compilation errors.
type UnsafeAuthorServer interface {
	mustEmbedUnimplementedAuthorServer()
}

func RegisterAuthorServer(s grpc.ServiceRegistrar, srv AuthorServer) {
	s.RegisterService(&Author_ServiceDesc, srv)
}

func _Author_GetByBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServer).GetByBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Author_GetByBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServer).GetByBook(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Author_ServiceDesc is the grpc.ServiceDesc for Author service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Author_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kvado.Author",
	HandlerType: (*AuthorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getByBook",
			Handler:    _Author_GetByBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "author.proto",
}