// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pkg/protobuf/urlspb/urlspb.proto

package urlspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	apipb "urlshortener/pkg/protobuf/apipb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UrlsClient is the client API for Urls service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlsClient interface {
	GetUrl(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error)
	CreateUrl(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlResponse, error)
	UpdateUrl(ctx context.Context, in *UpdateUrlRequest, opts ...grpc.CallOption) (*apipb.NoContent, error)
	DeleteUrl(ctx context.Context, in *DeleteUrlRequest, opts ...grpc.CallOption) (*apipb.NoContent, error)
	Ping(ctx context.Context, in *apipb.PingRequest, opts ...grpc.CallOption) (*apipb.PingResponse, error)
}

type urlsClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlsClient(cc grpc.ClientConnInterface) UrlsClient {
	return &urlsClient{cc}
}

func (c *urlsClient) GetUrl(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error) {
	out := new(GetUrlResponse)
	err := c.cc.Invoke(ctx, "/urlspb.Urls/GetUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlsClient) CreateUrl(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlResponse, error) {
	out := new(CreateUrlResponse)
	err := c.cc.Invoke(ctx, "/urlspb.Urls/CreateUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlsClient) UpdateUrl(ctx context.Context, in *UpdateUrlRequest, opts ...grpc.CallOption) (*apipb.NoContent, error) {
	out := new(apipb.NoContent)
	err := c.cc.Invoke(ctx, "/urlspb.Urls/UpdateUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlsClient) DeleteUrl(ctx context.Context, in *DeleteUrlRequest, opts ...grpc.CallOption) (*apipb.NoContent, error) {
	out := new(apipb.NoContent)
	err := c.cc.Invoke(ctx, "/urlspb.Urls/DeleteUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlsClient) Ping(ctx context.Context, in *apipb.PingRequest, opts ...grpc.CallOption) (*apipb.PingResponse, error) {
	out := new(apipb.PingResponse)
	err := c.cc.Invoke(ctx, "/urlspb.Urls/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlsServer is the server API for Urls service.
// All implementations must embed UnimplementedUrlsServer
// for forward compatibility
type UrlsServer interface {
	GetUrl(context.Context, *GetUrlRequest) (*GetUrlResponse, error)
	CreateUrl(context.Context, *CreateUrlRequest) (*CreateUrlResponse, error)
	UpdateUrl(context.Context, *UpdateUrlRequest) (*apipb.NoContent, error)
	DeleteUrl(context.Context, *DeleteUrlRequest) (*apipb.NoContent, error)
	Ping(context.Context, *apipb.PingRequest) (*apipb.PingResponse, error)
	mustEmbedUnimplementedUrlsServer()
}

// UnimplementedUrlsServer must be embedded to have forward compatible implementations.
type UnimplementedUrlsServer struct {
}

func (UnimplementedUrlsServer) GetUrl(context.Context, *GetUrlRequest) (*GetUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrl not implemented")
}
func (UnimplementedUrlsServer) CreateUrl(context.Context, *CreateUrlRequest) (*CreateUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUrl not implemented")
}
func (UnimplementedUrlsServer) UpdateUrl(context.Context, *UpdateUrlRequest) (*apipb.NoContent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUrl not implemented")
}
func (UnimplementedUrlsServer) DeleteUrl(context.Context, *DeleteUrlRequest) (*apipb.NoContent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUrl not implemented")
}
func (UnimplementedUrlsServer) Ping(context.Context, *apipb.PingRequest) (*apipb.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUrlsServer) mustEmbedUnimplementedUrlsServer() {}

// UnsafeUrlsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlsServer will
// result in compilation errors.
type UnsafeUrlsServer interface {
	mustEmbedUnimplementedUrlsServer()
}

func RegisterUrlsServer(s grpc.ServiceRegistrar, srv UrlsServer) {
	s.RegisterService(&Urls_ServiceDesc, srv)
}

func _Urls_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlspb.Urls/GetUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServer).GetUrl(ctx, req.(*GetUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Urls_CreateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServer).CreateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlspb.Urls/CreateUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServer).CreateUrl(ctx, req.(*CreateUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Urls_UpdateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServer).UpdateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlspb.Urls/UpdateUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServer).UpdateUrl(ctx, req.(*UpdateUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Urls_DeleteUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServer).DeleteUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlspb.Urls/DeleteUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServer).DeleteUrl(ctx, req.(*DeleteUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Urls_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(apipb.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlspb.Urls/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServer).Ping(ctx, req.(*apipb.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Urls_ServiceDesc is the grpc.ServiceDesc for Urls service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Urls_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "urlspb.Urls",
	HandlerType: (*UrlsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUrl",
			Handler:    _Urls_GetUrl_Handler,
		},
		{
			MethodName: "CreateUrl",
			Handler:    _Urls_CreateUrl_Handler,
		},
		{
			MethodName: "UpdateUrl",
			Handler:    _Urls_UpdateUrl_Handler,
		},
		{
			MethodName: "DeleteUrl",
			Handler:    _Urls_DeleteUrl_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Urls_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protobuf/urlspb/urlspb.proto",
}
