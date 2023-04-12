// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: book.proto

package book

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
	Book_GetByAuthor_FullMethodName = "/kvado.Book/getByAuthor"
)

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookClient interface {
	GetByAuthor(ctx context.Context, in *AuthorRequest, opts ...grpc.CallOption) (*BookListResponse, error)
}

type bookClient struct {
	cc grpc.ClientConnInterface
}

func NewBookClient(cc grpc.ClientConnInterface) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) GetByAuthor(ctx context.Context, in *AuthorRequest, opts ...grpc.CallOption) (*BookListResponse, error) {
	out := new(BookListResponse)
	err := c.cc.Invoke(ctx, Book_GetByAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServer is the server API for Book service.
// All implementations must embed UnimplementedBookServer
// for forward compatibility
type BookServer interface {
	GetByAuthor(context.Context, *AuthorRequest) (*BookListResponse, error)
	mustEmbedUnimplementedBookServer()
}

// UnimplementedBookServer must be embedded to have forward compatible implementations.
type UnimplementedBookServer struct {
}

func (UnimplementedBookServer) GetByAuthor(context.Context, *AuthorRequest) (*BookListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAuthor not implemented")
}
func (UnimplementedBookServer) mustEmbedUnimplementedBookServer() {}

// UnsafeBookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServer will
// result in compilation errors.
type UnsafeBookServer interface {
	mustEmbedUnimplementedBookServer()
}

func RegisterBookServer(s grpc.ServiceRegistrar, srv BookServer) {
	s.RegisterService(&Book_ServiceDesc, srv)
}

func _Book_GetByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetByAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetByAuthor(ctx, req.(*AuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Book_ServiceDesc is the grpc.ServiceDesc for Book service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Book_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kvado.Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getByAuthor",
			Handler:    _Book_GetByAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}