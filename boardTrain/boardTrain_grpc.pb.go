// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package boardTrain

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

// BoardingClient is the client API for Boarding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardingClient interface {
	PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketResponse, error)
	GetReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TicketResponse, error)
	RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	ViewSeats(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SeatResponse, error)
}

type boardingClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardingClient(cc grpc.ClientConnInterface) BoardingClient {
	return &boardingClient{cc}
}

func (c *boardingClient) PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketResponse, error) {
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, "/Boarding/purchaseTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardingClient) GetReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TicketResponse, error) {
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, "/Boarding/getReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardingClient) RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/Boarding/removeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardingClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/Boarding/modifySeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardingClient) ViewSeats(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SeatResponse, error) {
	out := new(SeatResponse)
	err := c.cc.Invoke(ctx, "/Boarding/viewSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardingServer is the server API for Boarding service.
// All implementations must embed UnimplementedBoardingServer
// for forward compatibility
type BoardingServer interface {
	PurchaseTicket(context.Context, *TicketRequest) (*TicketResponse, error)
	GetReceipt(context.Context, *UserRequest) (*TicketResponse, error)
	RemoveUser(context.Context, *UserRequest) (*StatusResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*StatusResponse, error)
	ViewSeats(context.Context, *SectionRequest) (*SeatResponse, error)
	mustEmbedUnimplementedBoardingServer()
}

// UnimplementedBoardingServer must be embedded to have forward compatible implementations.
type UnimplementedBoardingServer struct {
}

func (UnimplementedBoardingServer) PurchaseTicket(context.Context, *TicketRequest) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedBoardingServer) GetReceipt(context.Context, *UserRequest) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedBoardingServer) RemoveUser(context.Context, *UserRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedBoardingServer) ModifySeat(context.Context, *ModifySeatRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedBoardingServer) ViewSeats(context.Context, *SectionRequest) (*SeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewSeats not implemented")
}
func (UnimplementedBoardingServer) mustEmbedUnimplementedBoardingServer() {}

// UnsafeBoardingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardingServer will
// result in compilation errors.
type UnsafeBoardingServer interface {
	mustEmbedUnimplementedBoardingServer()
}

func RegisterBoardingServer(s grpc.ServiceRegistrar, srv BoardingServer) {
	s.RegisterService(&Boarding_ServiceDesc, srv)
}

func _Boarding_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardingServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Boarding/purchaseTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardingServer).PurchaseTicket(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boarding_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardingServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Boarding/getReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardingServer).GetReceipt(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boarding_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardingServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Boarding/removeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardingServer).RemoveUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boarding_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardingServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Boarding/modifySeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardingServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boarding_ViewSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardingServer).ViewSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Boarding/viewSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardingServer).ViewSeats(ctx, req.(*SectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Boarding_ServiceDesc is the grpc.ServiceDesc for Boarding service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Boarding_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Boarding",
	HandlerType: (*BoardingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "purchaseTicket",
			Handler:    _Boarding_PurchaseTicket_Handler,
		},
		{
			MethodName: "getReceipt",
			Handler:    _Boarding_GetReceipt_Handler,
		},
		{
			MethodName: "removeUser",
			Handler:    _Boarding_RemoveUser_Handler,
		},
		{
			MethodName: "modifySeat",
			Handler:    _Boarding_ModifySeat_Handler,
		},
		{
			MethodName: "viewSeats",
			Handler:    _Boarding_ViewSeats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "boardTrain.proto",
}
