// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: counter.proto

package pb

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

// CounterClient is the client API for Counter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterClient interface {
}

type counterClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterClient(cc grpc.ClientConnInterface) CounterClient {
	return &counterClient{cc}
}

// CounterServer is the server API for Counter service.
// All implementations must embed UnimplementedCounterServer
// for forward compatibility.
type CounterServer interface {
	mustEmbedUnimplementedCounterServer()
}

// UnimplementedCounterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCounterServer struct{}

func (UnimplementedCounterServer) mustEmbedUnimplementedCounterServer() {}
func (UnimplementedCounterServer) testEmbeddedByValue()                 {}

// UnsafeCounterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServer will
// result in compilation errors.
type UnsafeCounterServer interface {
	mustEmbedUnimplementedCounterServer()
}

func RegisterCounterServer(s grpc.ServiceRegistrar, srv CounterServer) {
	// If the following call pancis, it indicates UnimplementedCounterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Counter_ServiceDesc, srv)
}

// Counter_ServiceDesc is the grpc.ServiceDesc for Counter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Counter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.counter",
	HandlerType: (*CounterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "counter.proto",
}