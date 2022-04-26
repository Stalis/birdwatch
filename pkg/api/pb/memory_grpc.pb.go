// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: memory.proto

package pb

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

// MemoryClient is the client API for Memory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemoryClient interface {
	GetCurrentMemoryStats(ctx context.Context, in *CurrentMemoryRequest, opts ...grpc.CallOption) (*CurrentMemoryResponse, error)
	GetMemoryStats(ctx context.Context, in *MemoryStatsRequest, opts ...grpc.CallOption) (Memory_GetMemoryStatsClient, error)
}

type memoryClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoryClient(cc grpc.ClientConnInterface) MemoryClient {
	return &memoryClient{cc}
}

func (c *memoryClient) GetCurrentMemoryStats(ctx context.Context, in *CurrentMemoryRequest, opts ...grpc.CallOption) (*CurrentMemoryResponse, error) {
	out := new(CurrentMemoryResponse)
	err := c.cc.Invoke(ctx, "/memory.Memory/GetCurrentMemoryStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryClient) GetMemoryStats(ctx context.Context, in *MemoryStatsRequest, opts ...grpc.CallOption) (Memory_GetMemoryStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Memory_ServiceDesc.Streams[0], "/memory.Memory/GetMemoryStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &memoryGetMemoryStatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Memory_GetMemoryStatsClient interface {
	Recv() (*CurrentMemoryResponse, error)
	grpc.ClientStream
}

type memoryGetMemoryStatsClient struct {
	grpc.ClientStream
}

func (x *memoryGetMemoryStatsClient) Recv() (*CurrentMemoryResponse, error) {
	m := new(CurrentMemoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MemoryServer is the server API for Memory service.
// All implementations must embed UnimplementedMemoryServer
// for forward compatibility
type MemoryServer interface {
	GetCurrentMemoryStats(context.Context, *CurrentMemoryRequest) (*CurrentMemoryResponse, error)
	GetMemoryStats(*MemoryStatsRequest, Memory_GetMemoryStatsServer) error
	mustEmbedUnimplementedMemoryServer()
}

// UnimplementedMemoryServer must be embedded to have forward compatible implementations.
type UnimplementedMemoryServer struct {
}

func (UnimplementedMemoryServer) GetCurrentMemoryStats(context.Context, *CurrentMemoryRequest) (*CurrentMemoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentMemoryStats not implemented")
}
func (UnimplementedMemoryServer) GetMemoryStats(*MemoryStatsRequest, Memory_GetMemoryStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMemoryStats not implemented")
}
func (UnimplementedMemoryServer) mustEmbedUnimplementedMemoryServer() {}

// UnsafeMemoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemoryServer will
// result in compilation errors.
type UnsafeMemoryServer interface {
	mustEmbedUnimplementedMemoryServer()
}

func RegisterMemoryServer(s grpc.ServiceRegistrar, srv MemoryServer) {
	s.RegisterService(&Memory_ServiceDesc, srv)
}

func _Memory_GetCurrentMemoryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServer).GetCurrentMemoryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memory.Memory/GetCurrentMemoryStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServer).GetCurrentMemoryStats(ctx, req.(*CurrentMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memory_GetMemoryStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MemoryStatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MemoryServer).GetMemoryStats(m, &memoryGetMemoryStatsServer{stream})
}

type Memory_GetMemoryStatsServer interface {
	Send(*CurrentMemoryResponse) error
	grpc.ServerStream
}

type memoryGetMemoryStatsServer struct {
	grpc.ServerStream
}

func (x *memoryGetMemoryStatsServer) Send(m *CurrentMemoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Memory_ServiceDesc is the grpc.ServiceDesc for Memory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Memory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memory.Memory",
	HandlerType: (*MemoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentMemoryStats",
			Handler:    _Memory_GetCurrentMemoryStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMemoryStats",
			Handler:       _Memory_GetMemoryStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "memory.proto",
}
