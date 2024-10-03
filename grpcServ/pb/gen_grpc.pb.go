// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/gen.proto

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

// MoneyServiceClient is the client API for MoneyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoneyServiceClient interface {
	GetMoneyValues(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (*MoneyResponse, error)
}

type moneyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMoneyServiceClient(cc grpc.ClientConnInterface) MoneyServiceClient {
	return &moneyServiceClient{cc}
}

func (c *moneyServiceClient) GetMoneyValues(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (*MoneyResponse, error) {
	out := new(MoneyResponse)
	err := c.cc.Invoke(ctx, "/proto.MoneyService/GetMoneyValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoneyServiceServer is the cmd API for MoneyService service.
// All implementations must embed UnimplementedMoneyServiceServer
// for forward compatibility
type MoneyServiceServer interface {
	GetMoneyValues(context.Context, *MoneyRequest) (*MoneyResponse, error)
	mustEmbedUnimplementedMoneyServiceServer()
}

// UnimplementedMoneyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMoneyServiceServer struct {
}

func (UnimplementedMoneyServiceServer) GetMoneyValues(context.Context, *MoneyRequest) (*MoneyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoneyValues not implemented")
}
func (UnimplementedMoneyServiceServer) mustEmbedUnimplementedMoneyServiceServer() {}

// UnsafeMoneyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoneyServiceServer will
// result in compilation errors.
type UnsafeMoneyServiceServer interface {
	mustEmbedUnimplementedMoneyServiceServer()
}

func RegisterMoneyServiceServer(s grpc.ServiceRegistrar, srv MoneyServiceServer) {
	s.RegisterService(&MoneyService_ServiceDesc, srv)
}

func _MoneyService_GetMoneyValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoneyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).GetMoneyValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MoneyService/GetMoneyValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).GetMoneyValues(ctx, req.(*MoneyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MoneyService_ServiceDesc is the grpc.ServiceDesc for MoneyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoneyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MoneyService",
	HandlerType: (*MoneyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMoneyValues",
			Handler:    _MoneyService_GetMoneyValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gen.proto",
}
