// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: lemonade.proto

package lemonade_game

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

// LemonadeGameClient is the client API for LemonadeGame service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LemonadeGameClient interface {
	Create(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*CreateResult, error)
	RandomWeather(ctx context.Context, in *GameID, opts ...grpc.CallOption) (*Weather, error)
	GetBalance(ctx context.Context, in *GameID, opts ...grpc.CallOption) (*Balance, error)
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type lemonadeGameClient struct {
	cc grpc.ClientConnInterface
}

func NewLemonadeGameClient(cc grpc.ClientConnInterface) LemonadeGameClient {
	return &lemonadeGameClient{cc}
}

func (c *lemonadeGameClient) Create(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*CreateResult, error) {
	out := new(CreateResult)
	err := c.cc.Invoke(ctx, "/lemonade_game.LemonadeGame/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lemonadeGameClient) RandomWeather(ctx context.Context, in *GameID, opts ...grpc.CallOption) (*Weather, error) {
	out := new(Weather)
	err := c.cc.Invoke(ctx, "/lemonade_game.LemonadeGame/RandomWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lemonadeGameClient) GetBalance(ctx context.Context, in *GameID, opts ...grpc.CallOption) (*Balance, error) {
	out := new(Balance)
	err := c.cc.Invoke(ctx, "/lemonade_game.LemonadeGame/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lemonadeGameClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/lemonade_game.LemonadeGame/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LemonadeGameServer is the server API for LemonadeGame service.
// All implementations should embed UnimplementedLemonadeGameServer
// for forward compatibility
type LemonadeGameServer interface {
	Create(context.Context, *Nothing) (*CreateResult, error)
	RandomWeather(context.Context, *GameID) (*Weather, error)
	GetBalance(context.Context, *GameID) (*Balance, error)
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
}

// UnimplementedLemonadeGameServer should be embedded to have forward compatible implementations.
type UnimplementedLemonadeGameServer struct {
}

func (UnimplementedLemonadeGameServer) Create(context.Context, *Nothing) (*CreateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLemonadeGameServer) RandomWeather(context.Context, *GameID) (*Weather, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RandomWeather not implemented")
}
func (UnimplementedLemonadeGameServer) GetBalance(context.Context, *GameID) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedLemonadeGameServer) Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}

// UnsafeLemonadeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LemonadeGameServer will
// result in compilation errors.
type UnsafeLemonadeGameServer interface {
	mustEmbedUnimplementedLemonadeGameServer()
}

func RegisterLemonadeGameServer(s grpc.ServiceRegistrar, srv LemonadeGameServer) {
	s.RegisterService(&LemonadeGame_ServiceDesc, srv)
}

func _LemonadeGame_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LemonadeGameServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemonade_game.LemonadeGame/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LemonadeGameServer).Create(ctx, req.(*Nothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _LemonadeGame_RandomWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LemonadeGameServer).RandomWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemonade_game.LemonadeGame/RandomWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LemonadeGameServer).RandomWeather(ctx, req.(*GameID))
	}
	return interceptor(ctx, in, info, handler)
}

func _LemonadeGame_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LemonadeGameServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemonade_game.LemonadeGame/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LemonadeGameServer).GetBalance(ctx, req.(*GameID))
	}
	return interceptor(ctx, in, info, handler)
}

func _LemonadeGame_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LemonadeGameServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemonade_game.LemonadeGame/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LemonadeGameServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LemonadeGame_ServiceDesc is the grpc.ServiceDesc for LemonadeGame service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LemonadeGame_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lemonade_game.LemonadeGame",
	HandlerType: (*LemonadeGameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LemonadeGame_Create_Handler,
		},
		{
			MethodName: "RandomWeather",
			Handler:    _LemonadeGame_RandomWeather_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _LemonadeGame_GetBalance_Handler,
		},
		{
			MethodName: "Calculate",
			Handler:    _LemonadeGame_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lemonade.proto",
}