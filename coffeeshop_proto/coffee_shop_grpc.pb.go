// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: coffee_shop.proto

package golang_coffee_grpc

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

// CoffeeShopClient is the client API for CoffeeShop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoffeeShopClient interface {
	GetMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (CoffeeShop_GetMenuClient, error)
	PlaceOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Receipt, error)
	GetOrderStatus(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*OrderStatus, error)
}

type coffeeShopClient struct {
	cc grpc.ClientConnInterface
}

func NewCoffeeShopClient(cc grpc.ClientConnInterface) CoffeeShopClient {
	return &coffeeShopClient{cc}
}

func (c *coffeeShopClient) GetMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (CoffeeShop_GetMenuClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoffeeShop_ServiceDesc.Streams[0], "/coffeeshop.CoffeeShop/GetMenu", opts...)
	if err != nil {
		return nil, err
	}
	x := &coffeeShopGetMenuClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CoffeeShop_GetMenuClient interface {
	Recv() (*Menu, error)
	grpc.ClientStream
}

type coffeeShopGetMenuClient struct {
	grpc.ClientStream
}

func (x *coffeeShopGetMenuClient) Recv() (*Menu, error) {
	m := new(Menu)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coffeeShopClient) PlaceOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/coffeeshop.CoffeeShop/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coffeeShopClient) GetOrderStatus(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*OrderStatus, error) {
	out := new(OrderStatus)
	err := c.cc.Invoke(ctx, "/coffeeshop.CoffeeShop/GetOrderStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoffeeShopServer is the server API for CoffeeShop service.
// All implementations must embed UnimplementedCoffeeShopServer
// for forward compatibility
type CoffeeShopServer interface {
	GetMenu(*MenuRequest, CoffeeShop_GetMenuServer) error
	PlaceOrder(context.Context, *Order) (*Receipt, error)
	GetOrderStatus(context.Context, *Receipt) (*OrderStatus, error)
	mustEmbedUnimplementedCoffeeShopServer()
}

// UnimplementedCoffeeShopServer must be embedded to have forward compatible implementations.
type UnimplementedCoffeeShopServer struct {
}

func (UnimplementedCoffeeShopServer) GetMenu(*MenuRequest, CoffeeShop_GetMenuServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (UnimplementedCoffeeShopServer) PlaceOrder(context.Context, *Order) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedCoffeeShopServer) GetOrderStatus(context.Context, *Receipt) (*OrderStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStatus not implemented")
}
func (UnimplementedCoffeeShopServer) mustEmbedUnimplementedCoffeeShopServer() {}

// UnsafeCoffeeShopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoffeeShopServer will
// result in compilation errors.
type UnsafeCoffeeShopServer interface {
	mustEmbedUnimplementedCoffeeShopServer()
}

func RegisterCoffeeShopServer(s grpc.ServiceRegistrar, srv CoffeeShopServer) {
	s.RegisterService(&CoffeeShop_ServiceDesc, srv)
}

func _CoffeeShop_GetMenu_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MenuRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoffeeShopServer).GetMenu(m, &coffeeShopGetMenuServer{stream})
}

type CoffeeShop_GetMenuServer interface {
	Send(*Menu) error
	grpc.ServerStream
}

type coffeeShopGetMenuServer struct {
	grpc.ServerStream
}

func (x *coffeeShopGetMenuServer) Send(m *Menu) error {
	return x.ServerStream.SendMsg(m)
}

func _CoffeeShop_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeShopServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coffeeshop.CoffeeShop/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeShopServer).PlaceOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoffeeShop_GetOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Receipt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeShopServer).GetOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coffeeshop.CoffeeShop/GetOrderStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeShopServer).GetOrderStatus(ctx, req.(*Receipt))
	}
	return interceptor(ctx, in, info, handler)
}

// CoffeeShop_ServiceDesc is the grpc.ServiceDesc for CoffeeShop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoffeeShop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coffeeshop.CoffeeShop",
	HandlerType: (*CoffeeShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _CoffeeShop_PlaceOrder_Handler,
		},
		{
			MethodName: "GetOrderStatus",
			Handler:    _CoffeeShop_GetOrderStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMenu",
			Handler:       _CoffeeShop_GetMenu_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "coffee_shop.proto",
}
