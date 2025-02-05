// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: prime.proto

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

// PrimeCalculatorServiceClient is the client API for PrimeCalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrimeCalculatorServiceClient interface {
	Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeCalculatorService_PrimeClient, error)
}

type primeCalculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeCalculatorServiceClient(cc grpc.ClientConnInterface) PrimeCalculatorServiceClient {
	return &primeCalculatorServiceClient{cc}
}

func (c *primeCalculatorServiceClient) Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeCalculatorService_PrimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrimeCalculatorService_ServiceDesc.Streams[0], "/greet.PrimeCalculatorService/Prime", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeCalculatorServicePrimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeCalculatorService_PrimeClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type primeCalculatorServicePrimeClient struct {
	grpc.ClientStream
}

func (x *primeCalculatorServicePrimeClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeCalculatorServiceServer is the server API for PrimeCalculatorService service.
// All implementations must embed UnimplementedPrimeCalculatorServiceServer
// for forward compatibility
type PrimeCalculatorServiceServer interface {
	Prime(*PrimeRequest, PrimeCalculatorService_PrimeServer) error
	mustEmbedUnimplementedPrimeCalculatorServiceServer()
}

// UnimplementedPrimeCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrimeCalculatorServiceServer struct {
}

func (UnimplementedPrimeCalculatorServiceServer) Prime(*PrimeRequest, PrimeCalculatorService_PrimeServer) error {
	return status.Errorf(codes.Unimplemented, "method Prime not implemented")
}
func (UnimplementedPrimeCalculatorServiceServer) mustEmbedUnimplementedPrimeCalculatorServiceServer() {
}

// UnsafePrimeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrimeCalculatorServiceServer will
// result in compilation errors.
type UnsafePrimeCalculatorServiceServer interface {
	mustEmbedUnimplementedPrimeCalculatorServiceServer()
}

func RegisterPrimeCalculatorServiceServer(s grpc.ServiceRegistrar, srv PrimeCalculatorServiceServer) {
	s.RegisterService(&PrimeCalculatorService_ServiceDesc, srv)
}

func _PrimeCalculatorService_Prime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeCalculatorServiceServer).Prime(m, &primeCalculatorServicePrimeServer{stream})
}

type PrimeCalculatorService_PrimeServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type primeCalculatorServicePrimeServer struct {
	grpc.ServerStream
}

func (x *primeCalculatorServicePrimeServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PrimeCalculatorService_ServiceDesc is the grpc.ServiceDesc for PrimeCalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrimeCalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.PrimeCalculatorService",
	HandlerType: (*PrimeCalculatorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Prime",
			Handler:       _PrimeCalculatorService_Prime_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "prime.proto",
}
