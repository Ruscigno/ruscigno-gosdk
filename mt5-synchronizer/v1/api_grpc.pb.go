// Copyright 2021 Sander Ruscigno
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mql5_background_v1

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

// RuscignoMetatrader5ServiceClient is the client API for RuscignoMetatrader5Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuscignoMetatrader5ServiceClient interface {
	// Create a new account
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create a new deal
	CreateDeal(ctx context.Context, in *CreateDealRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create a new order
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create a new position
	CreatePosition(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type ruscignoMetatrader5ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRuscignoMetatrader5ServiceClient(cc grpc.ClientConnInterface) RuscignoMetatrader5ServiceClient {
	return &ruscignoMetatrader5ServiceClient{cc}
}

func (c *ruscignoMetatrader5ServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mql5_background.v1.RuscignoMetatrader5Service/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruscignoMetatrader5ServiceClient) CreateDeal(ctx context.Context, in *CreateDealRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mql5_background.v1.RuscignoMetatrader5Service/CreateDeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruscignoMetatrader5ServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mql5_background.v1.RuscignoMetatrader5Service/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruscignoMetatrader5ServiceClient) CreatePosition(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mql5_background.v1.RuscignoMetatrader5Service/CreatePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuscignoMetatrader5ServiceServer is the server API for RuscignoMetatrader5Service service.
// All implementations must embed UnimplementedRuscignoMetatrader5ServiceServer
// for forward compatibility
type RuscignoMetatrader5ServiceServer interface {
	// Create a new account
	CreateAccount(context.Context, *CreateAccountRequest) (*emptypb.Empty, error)
	// Create a new deal
	CreateDeal(context.Context, *CreateDealRequest) (*emptypb.Empty, error)
	// Create a new order
	CreateOrder(context.Context, *CreateOrderRequest) (*emptypb.Empty, error)
	// Create a new position
	CreatePosition(context.Context, *CreatePositionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRuscignoMetatrader5ServiceServer()
}

// UnimplementedRuscignoMetatrader5ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRuscignoMetatrader5ServiceServer struct {
}

func (UnimplementedRuscignoMetatrader5ServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedRuscignoMetatrader5ServiceServer) CreateDeal(context.Context, *CreateDealRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeal not implemented")
}
func (UnimplementedRuscignoMetatrader5ServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedRuscignoMetatrader5ServiceServer) CreatePosition(context.Context, *CreatePositionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePosition not implemented")
}
func (UnimplementedRuscignoMetatrader5ServiceServer) mustEmbedUnimplementedRuscignoMetatrader5ServiceServer() {
}

// UnsafeRuscignoMetatrader5ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuscignoMetatrader5ServiceServer will
// result in compilation errors.
type UnsafeRuscignoMetatrader5ServiceServer interface {
	mustEmbedUnimplementedRuscignoMetatrader5ServiceServer()
}

func RegisterRuscignoMetatrader5ServiceServer(s grpc.ServiceRegistrar, srv RuscignoMetatrader5ServiceServer) {
	s.RegisterService(&RuscignoMetatrader5Service_ServiceDesc, srv)
}

func _RuscignoMetatrader5Service_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuscignoMetatrader5ServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mql5_background.v1.RuscignoMetatrader5Service/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuscignoMetatrader5ServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuscignoMetatrader5Service_CreateDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuscignoMetatrader5ServiceServer).CreateDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mql5_background.v1.RuscignoMetatrader5Service/CreateDeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuscignoMetatrader5ServiceServer).CreateDeal(ctx, req.(*CreateDealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuscignoMetatrader5Service_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuscignoMetatrader5ServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mql5_background.v1.RuscignoMetatrader5Service/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuscignoMetatrader5ServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuscignoMetatrader5Service_CreatePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuscignoMetatrader5ServiceServer).CreatePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mql5_background.v1.RuscignoMetatrader5Service/CreatePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuscignoMetatrader5ServiceServer).CreatePosition(ctx, req.(*CreatePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RuscignoMetatrader5Service_ServiceDesc is the grpc.ServiceDesc for RuscignoMetatrader5Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuscignoMetatrader5Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mql5_background.v1.RuscignoMetatrader5Service",
	HandlerType: (*RuscignoMetatrader5ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _RuscignoMetatrader5Service_CreateAccount_Handler,
		},
		{
			MethodName: "CreateDeal",
			Handler:    _RuscignoMetatrader5Service_CreateDeal_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _RuscignoMetatrader5Service_CreateOrder_Handler,
		},
		{
			MethodName: "CreatePosition",
			Handler:    _RuscignoMetatrader5Service_CreatePosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
