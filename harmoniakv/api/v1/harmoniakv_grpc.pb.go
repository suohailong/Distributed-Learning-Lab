// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: api/v1/harmoniakv.proto

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

// HarmoniakvClient is the client API for Harmoniakv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HarmoniakvClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
}

type harmoniakvClient struct {
	cc grpc.ClientConnInterface
}

func NewHarmoniakvClient(cc grpc.ClientConnInterface) HarmoniakvClient {
	return &harmoniakvClient{cc}
}

func (c *harmoniakvClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/Harmoniakv/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harmoniakvClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/Harmoniakv/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HarmoniakvServer is the server API for Harmoniakv service.
// All implementations must embed UnimplementedHarmoniakvServer
// for forward compatibility
type HarmoniakvServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	mustEmbedUnimplementedHarmoniakvServer()
}

// UnimplementedHarmoniakvServer must be embedded to have forward compatible implementations.
type UnimplementedHarmoniakvServer struct {
}

func (UnimplementedHarmoniakvServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHarmoniakvServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedHarmoniakvServer) mustEmbedUnimplementedHarmoniakvServer() {}

// UnsafeHarmoniakvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HarmoniakvServer will
// result in compilation errors.
type UnsafeHarmoniakvServer interface {
	mustEmbedUnimplementedHarmoniakvServer()
}

func RegisterHarmoniakvServer(s grpc.ServiceRegistrar, srv HarmoniakvServer) {
	s.RegisterService(&Harmoniakv_ServiceDesc, srv)
}

func _Harmoniakv_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarmoniakvServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Harmoniakv/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarmoniakvServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harmoniakv_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarmoniakvServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Harmoniakv/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarmoniakvServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Harmoniakv_ServiceDesc is the grpc.ServiceDesc for Harmoniakv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Harmoniakv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Harmoniakv",
	HandlerType: (*HarmoniakvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Harmoniakv_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Harmoniakv_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/harmoniakv.proto",
}
