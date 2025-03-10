// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: mdb.proto

package proto

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
	MDBService_ListMachines_FullMethodName  = "/mdb.MDBService/ListMachines"
	MDBService_UpdateMachine_FullMethodName = "/mdb.MDBService/UpdateMachine"
	MDBService_GetMachine_FullMethodName    = "/mdb.MDBService/GetMachine"
)

// MDBServiceClient is the client API for MDBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MDBServiceClient interface {
	ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error)
	UpdateMachine(ctx context.Context, in *UpdateMachineRequest, opts ...grpc.CallOption) (*UpdateMachineResponse, error)
	GetMachine(ctx context.Context, in *GetMachineRequest, opts ...grpc.CallOption) (*GetMachineResponse, error)
}

type mDBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMDBServiceClient(cc grpc.ClientConnInterface) MDBServiceClient {
	return &mDBServiceClient{cc}
}

func (c *mDBServiceClient) ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error) {
	out := new(ListMachinesResponse)
	err := c.cc.Invoke(ctx, MDBService_ListMachines_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mDBServiceClient) UpdateMachine(ctx context.Context, in *UpdateMachineRequest, opts ...grpc.CallOption) (*UpdateMachineResponse, error) {
	out := new(UpdateMachineResponse)
	err := c.cc.Invoke(ctx, MDBService_UpdateMachine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mDBServiceClient) GetMachine(ctx context.Context, in *GetMachineRequest, opts ...grpc.CallOption) (*GetMachineResponse, error) {
	out := new(GetMachineResponse)
	err := c.cc.Invoke(ctx, MDBService_GetMachine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MDBServiceServer is the server API for MDBService service.
// All implementations should embed UnimplementedMDBServiceServer
// for forward compatibility
type MDBServiceServer interface {
	ListMachines(context.Context, *ListMachinesRequest) (*ListMachinesResponse, error)
	UpdateMachine(context.Context, *UpdateMachineRequest) (*UpdateMachineResponse, error)
	GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error)
}

// UnimplementedMDBServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMDBServiceServer struct {
}

func (UnimplementedMDBServiceServer) ListMachines(context.Context, *ListMachinesRequest) (*ListMachinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMachines not implemented")
}
func (UnimplementedMDBServiceServer) UpdateMachine(context.Context, *UpdateMachineRequest) (*UpdateMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMachine not implemented")
}
func (UnimplementedMDBServiceServer) GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachine not implemented")
}

// UnsafeMDBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MDBServiceServer will
// result in compilation errors.
type UnsafeMDBServiceServer interface {
	mustEmbedUnimplementedMDBServiceServer()
}

func RegisterMDBServiceServer(s grpc.ServiceRegistrar, srv MDBServiceServer) {
	s.RegisterService(&MDBService_ServiceDesc, srv)
}

func _MDBService_ListMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMachinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MDBServiceServer).ListMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MDBService_ListMachines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MDBServiceServer).ListMachines(ctx, req.(*ListMachinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MDBService_UpdateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MDBServiceServer).UpdateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MDBService_UpdateMachine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MDBServiceServer).UpdateMachine(ctx, req.(*UpdateMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MDBService_GetMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MDBServiceServer).GetMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MDBService_GetMachine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MDBServiceServer).GetMachine(ctx, req.(*GetMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MDBService_ServiceDesc is the grpc.ServiceDesc for MDBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MDBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mdb.MDBService",
	HandlerType: (*MDBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMachines",
			Handler:    _MDBService_ListMachines_Handler,
		},
		{
			MethodName: "UpdateMachine",
			Handler:    _MDBService_UpdateMachine_Handler,
		},
		{
			MethodName: "GetMachine",
			Handler:    _MDBService_GetMachine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mdb.proto",
}
