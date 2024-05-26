// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: encounter-service/encounter-service.proto

package encounter

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
	EncounterService_GetEncounter_FullMethodName    = "/EncounterService/GetEncounter"
	EncounterService_UpsertEncounter_FullMethodName = "/EncounterService/UpsertEncounter"
	EncounterService_DeleteEncounter_FullMethodName = "/EncounterService/DeleteEncounter"
)

// EncounterServiceClient is the client API for EncounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncounterServiceClient interface {
	GetEncounter(ctx context.Context, in *GetEncounterRequest, opts ...grpc.CallOption) (*GetEncounterResponse, error)
	UpsertEncounter(ctx context.Context, in *UpsertEncounterRequest, opts ...grpc.CallOption) (*UpsertEncounterResponse, error)
	DeleteEncounter(ctx context.Context, in *DeleteEncounterRequest, opts ...grpc.CallOption) (*DeleteEncounterResponse, error)
}

type encounterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncounterServiceClient(cc grpc.ClientConnInterface) EncounterServiceClient {
	return &encounterServiceClient{cc}
}

func (c *encounterServiceClient) GetEncounter(ctx context.Context, in *GetEncounterRequest, opts ...grpc.CallOption) (*GetEncounterResponse, error) {
	out := new(GetEncounterResponse)
	err := c.cc.Invoke(ctx, EncounterService_GetEncounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterServiceClient) UpsertEncounter(ctx context.Context, in *UpsertEncounterRequest, opts ...grpc.CallOption) (*UpsertEncounterResponse, error) {
	out := new(UpsertEncounterResponse)
	err := c.cc.Invoke(ctx, EncounterService_UpsertEncounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterServiceClient) DeleteEncounter(ctx context.Context, in *DeleteEncounterRequest, opts ...grpc.CallOption) (*DeleteEncounterResponse, error) {
	out := new(DeleteEncounterResponse)
	err := c.cc.Invoke(ctx, EncounterService_DeleteEncounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncounterServiceServer is the server API for EncounterService service.
// All implementations must embed UnimplementedEncounterServiceServer
// for forward compatibility
type EncounterServiceServer interface {
	GetEncounter(context.Context, *GetEncounterRequest) (*GetEncounterResponse, error)
	UpsertEncounter(context.Context, *UpsertEncounterRequest) (*UpsertEncounterResponse, error)
	DeleteEncounter(context.Context, *DeleteEncounterRequest) (*DeleteEncounterResponse, error)
	mustEmbedUnimplementedEncounterServiceServer()
}

// UnimplementedEncounterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEncounterServiceServer struct {
}

func (UnimplementedEncounterServiceServer) GetEncounter(context.Context, *GetEncounterRequest) (*GetEncounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncounter not implemented")
}
func (UnimplementedEncounterServiceServer) UpsertEncounter(context.Context, *UpsertEncounterRequest) (*UpsertEncounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertEncounter not implemented")
}
func (UnimplementedEncounterServiceServer) DeleteEncounter(context.Context, *DeleteEncounterRequest) (*DeleteEncounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEncounter not implemented")
}
func (UnimplementedEncounterServiceServer) mustEmbedUnimplementedEncounterServiceServer() {}

// UnsafeEncounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncounterServiceServer will
// result in compilation errors.
type UnsafeEncounterServiceServer interface {
	mustEmbedUnimplementedEncounterServiceServer()
}

func RegisterEncounterServiceServer(s grpc.ServiceRegistrar, srv EncounterServiceServer) {
	s.RegisterService(&EncounterService_ServiceDesc, srv)
}

func _EncounterService_GetEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEncounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServiceServer).GetEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterService_GetEncounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServiceServer).GetEncounter(ctx, req.(*GetEncounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncounterService_UpsertEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertEncounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServiceServer).UpsertEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterService_UpsertEncounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServiceServer).UpsertEncounter(ctx, req.(*UpsertEncounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncounterService_DeleteEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEncounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServiceServer).DeleteEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterService_DeleteEncounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServiceServer).DeleteEncounter(ctx, req.(*DeleteEncounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EncounterService_ServiceDesc is the grpc.ServiceDesc for EncounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EncounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EncounterService",
	HandlerType: (*EncounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEncounter",
			Handler:    _EncounterService_GetEncounter_Handler,
		},
		{
			MethodName: "UpsertEncounter",
			Handler:    _EncounterService_UpsertEncounter_Handler,
		},
		{
			MethodName: "DeleteEncounter",
			Handler:    _EncounterService_DeleteEncounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "encounter-service/encounter-service.proto",
}