// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: avg.proto

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

// AvgCalculatorServiceClient is the client API for AvgCalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvgCalculatorServiceClient interface {
	Avg(ctx context.Context, opts ...grpc.CallOption) (AvgCalculatorService_AvgClient, error)
}

type avgCalculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAvgCalculatorServiceClient(cc grpc.ClientConnInterface) AvgCalculatorServiceClient {
	return &avgCalculatorServiceClient{cc}
}

func (c *avgCalculatorServiceClient) Avg(ctx context.Context, opts ...grpc.CallOption) (AvgCalculatorService_AvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &AvgCalculatorService_ServiceDesc.Streams[0], "/avg.AvgCalculatorService/Avg", opts...)
	if err != nil {
		return nil, err
	}
	x := &avgCalculatorServiceAvgClient{stream}
	return x, nil
}

type AvgCalculatorService_AvgClient interface {
	Send(*AvgRequest) error
	CloseAndRecv() (*AvgResponse, error)
	grpc.ClientStream
}

type avgCalculatorServiceAvgClient struct {
	grpc.ClientStream
}

func (x *avgCalculatorServiceAvgClient) Send(m *AvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *avgCalculatorServiceAvgClient) CloseAndRecv() (*AvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AvgCalculatorServiceServer is the server API for AvgCalculatorService service.
// All implementations must embed UnimplementedAvgCalculatorServiceServer
// for forward compatibility
type AvgCalculatorServiceServer interface {
	Avg(AvgCalculatorService_AvgServer) error
	mustEmbedUnimplementedAvgCalculatorServiceServer()
}

// UnimplementedAvgCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAvgCalculatorServiceServer struct {
}

func (UnimplementedAvgCalculatorServiceServer) Avg(AvgCalculatorService_AvgServer) error {
	return status.Errorf(codes.Unimplemented, "method Avg not implemented")
}
func (UnimplementedAvgCalculatorServiceServer) mustEmbedUnimplementedAvgCalculatorServiceServer() {}

// UnsafeAvgCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvgCalculatorServiceServer will
// result in compilation errors.
type UnsafeAvgCalculatorServiceServer interface {
	mustEmbedUnimplementedAvgCalculatorServiceServer()
}

func RegisterAvgCalculatorServiceServer(s grpc.ServiceRegistrar, srv AvgCalculatorServiceServer) {
	s.RegisterService(&AvgCalculatorService_ServiceDesc, srv)
}

func _AvgCalculatorService_Avg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AvgCalculatorServiceServer).Avg(&avgCalculatorServiceAvgServer{stream})
}

type AvgCalculatorService_AvgServer interface {
	SendAndClose(*AvgResponse) error
	Recv() (*AvgRequest, error)
	grpc.ServerStream
}

type avgCalculatorServiceAvgServer struct {
	grpc.ServerStream
}

func (x *avgCalculatorServiceAvgServer) SendAndClose(m *AvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *avgCalculatorServiceAvgServer) Recv() (*AvgRequest, error) {
	m := new(AvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AvgCalculatorService_ServiceDesc is the grpc.ServiceDesc for AvgCalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AvgCalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "avg.AvgCalculatorService",
	HandlerType: (*AvgCalculatorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Avg",
			Handler:       _AvgCalculatorService_Avg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "avg.proto",
}
