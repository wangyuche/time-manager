// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: timemanager/timemanager.proto

package timemanager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TimeGRPC_GetServerTime_FullMethodName       = "/timemanager.TimeGRPC/GetServerTime"
	TimeGRPC_SetServerTime_FullMethodName       = "/timemanager.TimeGRPC/SetServerTime"
	TimeGRPC_ClearServerTime_FullMethodName     = "/timemanager.TimeGRPC/ClearServerTime"
	TimeGRPC_BroadcastServerTime_FullMethodName = "/timemanager.TimeGRPC/BroadcastServerTime"
)

// TimeGRPCClient is the client API for TimeGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeGRPCClient interface {
	GetServerTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetServerTimeRes, error)
	SetServerTime(ctx context.Context, in *SetServerTimeReq, opts ...grpc.CallOption) (*SetServerTimeRes, error)
	ClearServerTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClearServerTimeRes, error)
	BroadcastServerTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TimeGRPC_BroadcastServerTimeClient, error)
}

type timeGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeGRPCClient(cc grpc.ClientConnInterface) TimeGRPCClient {
	return &timeGRPCClient{cc}
}

func (c *timeGRPCClient) GetServerTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetServerTimeRes, error) {
	out := new(GetServerTimeRes)
	err := c.cc.Invoke(ctx, TimeGRPC_GetServerTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeGRPCClient) SetServerTime(ctx context.Context, in *SetServerTimeReq, opts ...grpc.CallOption) (*SetServerTimeRes, error) {
	out := new(SetServerTimeRes)
	err := c.cc.Invoke(ctx, TimeGRPC_SetServerTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeGRPCClient) ClearServerTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClearServerTimeRes, error) {
	out := new(ClearServerTimeRes)
	err := c.cc.Invoke(ctx, TimeGRPC_ClearServerTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeGRPCClient) BroadcastServerTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TimeGRPC_BroadcastServerTimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &TimeGRPC_ServiceDesc.Streams[0], TimeGRPC_BroadcastServerTime_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &timeGRPCBroadcastServerTimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TimeGRPC_BroadcastServerTimeClient interface {
	Recv() (*GetServerTimeRes, error)
	grpc.ClientStream
}

type timeGRPCBroadcastServerTimeClient struct {
	grpc.ClientStream
}

func (x *timeGRPCBroadcastServerTimeClient) Recv() (*GetServerTimeRes, error) {
	m := new(GetServerTimeRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TimeGRPCServer is the server API for TimeGRPC service.
// All implementations must embed UnimplementedTimeGRPCServer
// for forward compatibility
type TimeGRPCServer interface {
	GetServerTime(context.Context, *emptypb.Empty) (*GetServerTimeRes, error)
	SetServerTime(context.Context, *SetServerTimeReq) (*SetServerTimeRes, error)
	ClearServerTime(context.Context, *emptypb.Empty) (*ClearServerTimeRes, error)
	BroadcastServerTime(*emptypb.Empty, TimeGRPC_BroadcastServerTimeServer) error
	mustEmbedUnimplementedTimeGRPCServer()
}

// UnimplementedTimeGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedTimeGRPCServer struct {
}

func (UnimplementedTimeGRPCServer) GetServerTime(context.Context, *emptypb.Empty) (*GetServerTimeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerTime not implemented")
}
func (UnimplementedTimeGRPCServer) SetServerTime(context.Context, *SetServerTimeReq) (*SetServerTimeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetServerTime not implemented")
}
func (UnimplementedTimeGRPCServer) ClearServerTime(context.Context, *emptypb.Empty) (*ClearServerTimeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearServerTime not implemented")
}
func (UnimplementedTimeGRPCServer) BroadcastServerTime(*emptypb.Empty, TimeGRPC_BroadcastServerTimeServer) error {
	return status.Errorf(codes.Unimplemented, "method BroadcastServerTime not implemented")
}
func (UnimplementedTimeGRPCServer) mustEmbedUnimplementedTimeGRPCServer() {}

// UnsafeTimeGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeGRPCServer will
// result in compilation errors.
type UnsafeTimeGRPCServer interface {
	mustEmbedUnimplementedTimeGRPCServer()
}

func RegisterTimeGRPCServer(s grpc.ServiceRegistrar, srv TimeGRPCServer) {
	s.RegisterService(&TimeGRPC_ServiceDesc, srv)
}

func _TimeGRPC_GetServerTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeGRPCServer).GetServerTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeGRPC_GetServerTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeGRPCServer).GetServerTime(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeGRPC_SetServerTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetServerTimeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeGRPCServer).SetServerTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeGRPC_SetServerTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeGRPCServer).SetServerTime(ctx, req.(*SetServerTimeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeGRPC_ClearServerTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeGRPCServer).ClearServerTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeGRPC_ClearServerTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeGRPCServer).ClearServerTime(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeGRPC_BroadcastServerTime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TimeGRPCServer).BroadcastServerTime(m, &timeGRPCBroadcastServerTimeServer{stream})
}

type TimeGRPC_BroadcastServerTimeServer interface {
	Send(*GetServerTimeRes) error
	grpc.ServerStream
}

type timeGRPCBroadcastServerTimeServer struct {
	grpc.ServerStream
}

func (x *timeGRPCBroadcastServerTimeServer) Send(m *GetServerTimeRes) error {
	return x.ServerStream.SendMsg(m)
}

// TimeGRPC_ServiceDesc is the grpc.ServiceDesc for TimeGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timemanager.TimeGRPC",
	HandlerType: (*TimeGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerTime",
			Handler:    _TimeGRPC_GetServerTime_Handler,
		},
		{
			MethodName: "SetServerTime",
			Handler:    _TimeGRPC_SetServerTime_Handler,
		},
		{
			MethodName: "ClearServerTime",
			Handler:    _TimeGRPC_ClearServerTime_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BroadcastServerTime",
			Handler:       _TimeGRPC_BroadcastServerTime_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "timemanager/timemanager.proto",
}
