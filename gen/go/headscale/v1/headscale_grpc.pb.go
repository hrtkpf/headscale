// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// HeadscaleServiceClient is the client API for HeadscaleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeadscaleServiceClient interface {
	GetMachine(ctx context.Context, in *GetMachineRequest, opts ...grpc.CallOption) (*GetMachineResponse, error)
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error)
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
}

type headscaleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeadscaleServiceClient(cc grpc.ClientConnInterface) HeadscaleServiceClient {
	return &headscaleServiceClient{cc}
}

func (c *headscaleServiceClient) GetMachine(ctx context.Context, in *GetMachineRequest, opts ...grpc.CallOption) (*GetMachineResponse, error) {
	out := new(GetMachineResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/GetMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error) {
	out := new(CreateNamespaceResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error) {
	out := new(DeleteNamespaceResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/ListNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeadscaleServiceServer is the server API for HeadscaleService service.
// All implementations must embed UnimplementedHeadscaleServiceServer
// for forward compatibility
type HeadscaleServiceServer interface {
	GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error)
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	mustEmbedUnimplementedHeadscaleServiceServer()
}

// UnimplementedHeadscaleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeadscaleServiceServer struct {
}

func (UnimplementedHeadscaleServiceServer) GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachine not implemented")
}
func (UnimplementedHeadscaleServiceServer) CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedHeadscaleServiceServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedHeadscaleServiceServer) ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}
func (UnimplementedHeadscaleServiceServer) mustEmbedUnimplementedHeadscaleServiceServer() {}

// UnsafeHeadscaleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeadscaleServiceServer will
// result in compilation errors.
type UnsafeHeadscaleServiceServer interface {
	mustEmbedUnimplementedHeadscaleServiceServer()
}

func RegisterHeadscaleServiceServer(s grpc.ServiceRegistrar, srv HeadscaleServiceServer) {
	s.RegisterService(&HeadscaleService_ServiceDesc, srv)
}

func _HeadscaleService_GetMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).GetMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/GetMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).GetMachine(ctx, req.(*GetMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HeadscaleService_ServiceDesc is the grpc.ServiceDesc for HeadscaleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeadscaleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "headscale.v1.HeadscaleService",
	HandlerType: (*HeadscaleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMachine",
			Handler:    _HeadscaleService_GetMachine_Handler,
		},
		{
			MethodName: "CreateNamespace",
			Handler:    _HeadscaleService_CreateNamespace_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _HeadscaleService_DeleteNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _HeadscaleService_ListNamespaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "headscale/v1/headscale.proto",
}