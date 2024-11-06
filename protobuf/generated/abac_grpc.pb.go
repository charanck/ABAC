// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: protobuf/abac.proto

package abac

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

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	Health(ctx context.Context, in *Void, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Health(ctx context.Context, in *Void, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/abac.Health/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
// All implementations must embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	Health(context.Context, *Void) (*HealthResponse, error)
	mustEmbedUnimplementedHealthServer()
}

// UnimplementedHealthServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) Health(context.Context, *Void) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedHealthServer) mustEmbedUnimplementedHealthServer() {}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&Health_ServiceDesc, srv)
}

func _Health_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Health/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Health(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// Health_ServiceDesc is the grpc.ServiceDesc for Health service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Health_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "abac.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Health_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/abac.proto",
}

// ResourceClient is the client API for Resource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceClient interface {
	CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceResponse, error)
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceResponse, error)
	UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceResponse, error)
	DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*DeleteResourceResponse, error)
	ListResource(ctx context.Context, in *ListResourceRequest, opts ...grpc.CallOption) (*ListResourceResponse, error)
}

type resourceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceClient(cc grpc.ClientConnInterface) ResourceClient {
	return &resourceClient{cc}
}

func (c *resourceClient) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceResponse, error) {
	out := new(CreateResourceResponse)
	err := c.cc.Invoke(ctx, "/abac.Resource/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceResponse, error) {
	out := new(GetResourceResponse)
	err := c.cc.Invoke(ctx, "/abac.Resource/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceResponse, error) {
	out := new(UpdateResourceResponse)
	err := c.cc.Invoke(ctx, "/abac.Resource/UpdateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*DeleteResourceResponse, error) {
	out := new(DeleteResourceResponse)
	err := c.cc.Invoke(ctx, "/abac.Resource/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) ListResource(ctx context.Context, in *ListResourceRequest, opts ...grpc.CallOption) (*ListResourceResponse, error) {
	out := new(ListResourceResponse)
	err := c.cc.Invoke(ctx, "/abac.Resource/ListResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServer is the server API for Resource service.
// All implementations must embed UnimplementedResourceServer
// for forward compatibility
type ResourceServer interface {
	CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error)
	GetResource(context.Context, *GetResourceRequest) (*GetResourceResponse, error)
	UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceResponse, error)
	DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error)
	ListResource(context.Context, *ListResourceRequest) (*ListResourceResponse, error)
	mustEmbedUnimplementedResourceServer()
}

// UnimplementedResourceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceServer struct {
}

func (UnimplementedResourceServer) CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedResourceServer) GetResource(context.Context, *GetResourceRequest) (*GetResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResource not implemented")
}
func (UnimplementedResourceServer) UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}
func (UnimplementedResourceServer) DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedResourceServer) ListResource(context.Context, *ListResourceRequest) (*ListResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResource not implemented")
}
func (UnimplementedResourceServer) mustEmbedUnimplementedResourceServer() {}

// UnsafeResourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServer will
// result in compilation errors.
type UnsafeResourceServer interface {
	mustEmbedUnimplementedResourceServer()
}

func RegisterResourceServer(s grpc.ServiceRegistrar, srv ResourceServer) {
	s.RegisterService(&Resource_ServiceDesc, srv)
}

func _Resource_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Resource/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).CreateResource(ctx, req.(*CreateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Resource/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).GetResource(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Resource/UpdateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).UpdateResource(ctx, req.(*UpdateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Resource/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).DeleteResource(ctx, req.(*DeleteResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_ListResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).ListResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Resource/ListResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).ListResource(ctx, req.(*ListResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Resource_ServiceDesc is the grpc.ServiceDesc for Resource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "abac.Resource",
	HandlerType: (*ResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _Resource_CreateResource_Handler,
		},
		{
			MethodName: "GetResource",
			Handler:    _Resource_GetResource_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _Resource_UpdateResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _Resource_DeleteResource_Handler,
		},
		{
			MethodName: "ListResource",
			Handler:    _Resource_ListResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/abac.proto",
}

// AttributeClient is the client API for Attribute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttributeClient interface {
	CreateAttribute(ctx context.Context, in *CreateAttributeRequest, opts ...grpc.CallOption) (*CreateAttributeResponse, error)
	BatchCreateAttribute(ctx context.Context, in *BatchCreateAttributeRequest, opts ...grpc.CallOption) (*BatchCreateAttributeResponse, error)
	UpdateAttribute(ctx context.Context, in *UpdateAttributeRequest, opts ...grpc.CallOption) (*UpdateAttributeResponse, error)
	ListAttribute(ctx context.Context, in *ListAttributeRequest, opts ...grpc.CallOption) (*ListAttributeResponse, error)
	GetAttribute(ctx context.Context, in *GetAttributeRequest, opts ...grpc.CallOption) (*GetAttributeResponse, error)
	DeleteAttribute(ctx context.Context, in *DeleteAttributeRequest, opts ...grpc.CallOption) (*Void, error)
}

type attributeClient struct {
	cc grpc.ClientConnInterface
}

func NewAttributeClient(cc grpc.ClientConnInterface) AttributeClient {
	return &attributeClient{cc}
}

func (c *attributeClient) CreateAttribute(ctx context.Context, in *CreateAttributeRequest, opts ...grpc.CallOption) (*CreateAttributeResponse, error) {
	out := new(CreateAttributeResponse)
	err := c.cc.Invoke(ctx, "/abac.Attribute/CreateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeClient) BatchCreateAttribute(ctx context.Context, in *BatchCreateAttributeRequest, opts ...grpc.CallOption) (*BatchCreateAttributeResponse, error) {
	out := new(BatchCreateAttributeResponse)
	err := c.cc.Invoke(ctx, "/abac.Attribute/BatchCreateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeClient) UpdateAttribute(ctx context.Context, in *UpdateAttributeRequest, opts ...grpc.CallOption) (*UpdateAttributeResponse, error) {
	out := new(UpdateAttributeResponse)
	err := c.cc.Invoke(ctx, "/abac.Attribute/UpdateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeClient) ListAttribute(ctx context.Context, in *ListAttributeRequest, opts ...grpc.CallOption) (*ListAttributeResponse, error) {
	out := new(ListAttributeResponse)
	err := c.cc.Invoke(ctx, "/abac.Attribute/ListAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeClient) GetAttribute(ctx context.Context, in *GetAttributeRequest, opts ...grpc.CallOption) (*GetAttributeResponse, error) {
	out := new(GetAttributeResponse)
	err := c.cc.Invoke(ctx, "/abac.Attribute/GetAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeClient) DeleteAttribute(ctx context.Context, in *DeleteAttributeRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/abac.Attribute/DeleteAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttributeServer is the server API for Attribute service.
// All implementations must embed UnimplementedAttributeServer
// for forward compatibility
type AttributeServer interface {
	CreateAttribute(context.Context, *CreateAttributeRequest) (*CreateAttributeResponse, error)
	BatchCreateAttribute(context.Context, *BatchCreateAttributeRequest) (*BatchCreateAttributeResponse, error)
	UpdateAttribute(context.Context, *UpdateAttributeRequest) (*UpdateAttributeResponse, error)
	ListAttribute(context.Context, *ListAttributeRequest) (*ListAttributeResponse, error)
	GetAttribute(context.Context, *GetAttributeRequest) (*GetAttributeResponse, error)
	DeleteAttribute(context.Context, *DeleteAttributeRequest) (*Void, error)
	mustEmbedUnimplementedAttributeServer()
}

// UnimplementedAttributeServer must be embedded to have forward compatible implementations.
type UnimplementedAttributeServer struct {
}

func (UnimplementedAttributeServer) CreateAttribute(context.Context, *CreateAttributeRequest) (*CreateAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttribute not implemented")
}
func (UnimplementedAttributeServer) BatchCreateAttribute(context.Context, *BatchCreateAttributeRequest) (*BatchCreateAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreateAttribute not implemented")
}
func (UnimplementedAttributeServer) UpdateAttribute(context.Context, *UpdateAttributeRequest) (*UpdateAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAttribute not implemented")
}
func (UnimplementedAttributeServer) ListAttribute(context.Context, *ListAttributeRequest) (*ListAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttribute not implemented")
}
func (UnimplementedAttributeServer) GetAttribute(context.Context, *GetAttributeRequest) (*GetAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttribute not implemented")
}
func (UnimplementedAttributeServer) DeleteAttribute(context.Context, *DeleteAttributeRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAttribute not implemented")
}
func (UnimplementedAttributeServer) mustEmbedUnimplementedAttributeServer() {}

// UnsafeAttributeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttributeServer will
// result in compilation errors.
type UnsafeAttributeServer interface {
	mustEmbedUnimplementedAttributeServer()
}

func RegisterAttributeServer(s grpc.ServiceRegistrar, srv AttributeServer) {
	s.RegisterService(&Attribute_ServiceDesc, srv)
}

func _Attribute_CreateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServer).CreateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Attribute/CreateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServer).CreateAttribute(ctx, req.(*CreateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attribute_BatchCreateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServer).BatchCreateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Attribute/BatchCreateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServer).BatchCreateAttribute(ctx, req.(*BatchCreateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attribute_UpdateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServer).UpdateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Attribute/UpdateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServer).UpdateAttribute(ctx, req.(*UpdateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attribute_ListAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServer).ListAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Attribute/ListAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServer).ListAttribute(ctx, req.(*ListAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attribute_GetAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServer).GetAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Attribute/GetAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServer).GetAttribute(ctx, req.(*GetAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attribute_DeleteAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServer).DeleteAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abac.Attribute/DeleteAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServer).DeleteAttribute(ctx, req.(*DeleteAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Attribute_ServiceDesc is the grpc.ServiceDesc for Attribute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Attribute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "abac.Attribute",
	HandlerType: (*AttributeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAttribute",
			Handler:    _Attribute_CreateAttribute_Handler,
		},
		{
			MethodName: "BatchCreateAttribute",
			Handler:    _Attribute_BatchCreateAttribute_Handler,
		},
		{
			MethodName: "UpdateAttribute",
			Handler:    _Attribute_UpdateAttribute_Handler,
		},
		{
			MethodName: "ListAttribute",
			Handler:    _Attribute_ListAttribute_Handler,
		},
		{
			MethodName: "GetAttribute",
			Handler:    _Attribute_GetAttribute_Handler,
		},
		{
			MethodName: "DeleteAttribute",
			Handler:    _Attribute_DeleteAttribute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/abac.proto",
}
