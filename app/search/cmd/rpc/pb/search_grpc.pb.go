// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: search.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Search_SearchUsers_FullMethodName   = "/pb.search/SearchUsers"
	Search_SearchContent_FullMethodName = "/pb.search/SearchContent"
)

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchClient interface {
	SearchUsers(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchUsersResp, error)
	SearchContent(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchContentResp, error)
}

type searchClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchClient(cc grpc.ClientConnInterface) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) SearchUsers(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchUsersResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchUsersResp)
	err := c.cc.Invoke(ctx, Search_SearchUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) SearchContent(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchContentResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchContentResp)
	err := c.cc.Invoke(ctx, Search_SearchContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServer is the server API for Search service.
// All implementations must embed UnimplementedSearchServer
// for forward compatibility.
type SearchServer interface {
	SearchUsers(context.Context, *SearchReq) (*SearchUsersResp, error)
	SearchContent(context.Context, *SearchReq) (*SearchContentResp, error)
	mustEmbedUnimplementedSearchServer()
}

// UnimplementedSearchServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchServer struct{}

func (UnimplementedSearchServer) SearchUsers(context.Context, *SearchReq) (*SearchUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
func (UnimplementedSearchServer) SearchContent(context.Context, *SearchReq) (*SearchContentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchContent not implemented")
}
func (UnimplementedSearchServer) mustEmbedUnimplementedSearchServer() {}
func (UnimplementedSearchServer) testEmbeddedByValue()                {}

// UnsafeSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServer will
// result in compilation errors.
type UnsafeSearchServer interface {
	mustEmbedUnimplementedSearchServer()
}

func RegisterSearchServer(s grpc.ServiceRegistrar, srv SearchServer) {
	// If the following call pancis, it indicates UnimplementedSearchServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Search_ServiceDesc, srv)
}

func _Search_SearchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).SearchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Search_SearchUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).SearchUsers(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_SearchContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).SearchContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Search_SearchContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).SearchContent(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Search_ServiceDesc is the grpc.ServiceDesc for Search service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Search_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUsers",
			Handler:    _Search_SearchUsers_Handler,
		},
		{
			MethodName: "SearchContent",
			Handler:    _Search_SearchContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}