// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: seo.proto

package protos

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
	SEO_GetSEO_FullMethodName    = "/protos.SEO/GetSEO"
	SEO_CreateSEO_FullMethodName = "/protos.SEO/CreateSEO"
	SEO_UpdateSEO_FullMethodName = "/protos.SEO/UpdateSEO"
	SEO_DeleteSEO_FullMethodName = "/protos.SEO/DeleteSEO"
)

// SEOClient is the client API for SEO service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SEOClient interface {
	GetSEO(ctx context.Context, in *GetSEOReq, opts ...grpc.CallOption) (*SEOMsg, error)
	CreateSEO(ctx context.Context, in *CreateAndUpdateSEOReq, opts ...grpc.CallOption) (*EmptySEO, error)
	UpdateSEO(ctx context.Context, in *CreateAndUpdateSEOReq, opts ...grpc.CallOption) (*EmptySEO, error)
	DeleteSEO(ctx context.Context, in *GetSEOReq, opts ...grpc.CallOption) (*EmptySEO, error)
}

type sEOClient struct {
	cc grpc.ClientConnInterface
}

func NewSEOClient(cc grpc.ClientConnInterface) SEOClient {
	return &sEOClient{cc}
}

func (c *sEOClient) GetSEO(ctx context.Context, in *GetSEOReq, opts ...grpc.CallOption) (*SEOMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SEOMsg)
	err := c.cc.Invoke(ctx, SEO_GetSEO_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sEOClient) CreateSEO(ctx context.Context, in *CreateAndUpdateSEOReq, opts ...grpc.CallOption) (*EmptySEO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptySEO)
	err := c.cc.Invoke(ctx, SEO_CreateSEO_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sEOClient) UpdateSEO(ctx context.Context, in *CreateAndUpdateSEOReq, opts ...grpc.CallOption) (*EmptySEO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptySEO)
	err := c.cc.Invoke(ctx, SEO_UpdateSEO_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sEOClient) DeleteSEO(ctx context.Context, in *GetSEOReq, opts ...grpc.CallOption) (*EmptySEO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptySEO)
	err := c.cc.Invoke(ctx, SEO_DeleteSEO_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SEOServer is the server API for SEO service.
// All implementations must embed UnimplementedSEOServer
// for forward compatibility.
type SEOServer interface {
	GetSEO(context.Context, *GetSEOReq) (*SEOMsg, error)
	CreateSEO(context.Context, *CreateAndUpdateSEOReq) (*EmptySEO, error)
	UpdateSEO(context.Context, *CreateAndUpdateSEOReq) (*EmptySEO, error)
	DeleteSEO(context.Context, *GetSEOReq) (*EmptySEO, error)
	mustEmbedUnimplementedSEOServer()
}

// UnimplementedSEOServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSEOServer struct{}

func (UnimplementedSEOServer) GetSEO(context.Context, *GetSEOReq) (*SEOMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSEO not implemented")
}
func (UnimplementedSEOServer) CreateSEO(context.Context, *CreateAndUpdateSEOReq) (*EmptySEO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSEO not implemented")
}
func (UnimplementedSEOServer) UpdateSEO(context.Context, *CreateAndUpdateSEOReq) (*EmptySEO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSEO not implemented")
}
func (UnimplementedSEOServer) DeleteSEO(context.Context, *GetSEOReq) (*EmptySEO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSEO not implemented")
}
func (UnimplementedSEOServer) mustEmbedUnimplementedSEOServer() {}
func (UnimplementedSEOServer) testEmbeddedByValue()             {}

// UnsafeSEOServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SEOServer will
// result in compilation errors.
type UnsafeSEOServer interface {
	mustEmbedUnimplementedSEOServer()
}

func RegisterSEOServer(s grpc.ServiceRegistrar, srv SEOServer) {
	// If the following call pancis, it indicates UnimplementedSEOServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SEO_ServiceDesc, srv)
}

func _SEO_GetSEO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSEOReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SEOServer).GetSEO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SEO_GetSEO_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SEOServer).GetSEO(ctx, req.(*GetSEOReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SEO_CreateSEO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAndUpdateSEOReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SEOServer).CreateSEO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SEO_CreateSEO_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SEOServer).CreateSEO(ctx, req.(*CreateAndUpdateSEOReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SEO_UpdateSEO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAndUpdateSEOReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SEOServer).UpdateSEO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SEO_UpdateSEO_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SEOServer).UpdateSEO(ctx, req.(*CreateAndUpdateSEOReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SEO_DeleteSEO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSEOReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SEOServer).DeleteSEO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SEO_DeleteSEO_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SEOServer).DeleteSEO(ctx, req.(*GetSEOReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SEO_ServiceDesc is the grpc.ServiceDesc for SEO service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SEO_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.SEO",
	HandlerType: (*SEOServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSEO",
			Handler:    _SEO_GetSEO_Handler,
		},
		{
			MethodName: "CreateSEO",
			Handler:    _SEO_CreateSEO_Handler,
		},
		{
			MethodName: "UpdateSEO",
			Handler:    _SEO_UpdateSEO_Handler,
		},
		{
			MethodName: "DeleteSEO",
			Handler:    _SEO_DeleteSEO_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seo.proto",
}
