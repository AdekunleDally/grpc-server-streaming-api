// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: pf_calculator.proto

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

// PrimeFactorServiceClient is the client API for PrimeFactorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrimeFactorServiceClient interface {
	PrimeFactorCalc(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*PrimeFactorResponse, error)
}

type primeFactorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeFactorServiceClient(cc grpc.ClientConnInterface) PrimeFactorServiceClient {
	return &primeFactorServiceClient{cc}
}

func (c *primeFactorServiceClient) PrimeFactorCalc(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*PrimeFactorResponse, error) {
	out := new(PrimeFactorResponse)
	err := c.cc.Invoke(ctx, "/primefactorcalculator.PrimeFactorService/PrimeFactorCalc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrimeFactorServiceServer is the server API for PrimeFactorService service.
// All implementations must embed UnimplementedPrimeFactorServiceServer
// for forward compatibility
type PrimeFactorServiceServer interface {
	PrimeFactorCalc(context.Context, *NumberRequest) (*PrimeFactorResponse, error)
	mustEmbedUnimplementedPrimeFactorServiceServer()
}

// UnimplementedPrimeFactorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrimeFactorServiceServer struct {
}

func (UnimplementedPrimeFactorServiceServer) PrimeFactorCalc(context.Context, *NumberRequest) (*PrimeFactorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrimeFactorCalc not implemented")
}
func (UnimplementedPrimeFactorServiceServer) mustEmbedUnimplementedPrimeFactorServiceServer() {}

// UnsafePrimeFactorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrimeFactorServiceServer will
// result in compilation errors.
type UnsafePrimeFactorServiceServer interface {
	mustEmbedUnimplementedPrimeFactorServiceServer()
}

func RegisterPrimeFactorServiceServer(s grpc.ServiceRegistrar, srv PrimeFactorServiceServer) {
	s.RegisterService(&PrimeFactorService_ServiceDesc, srv)
}

func _PrimeFactorService_PrimeFactorCalc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrimeFactorServiceServer).PrimeFactorCalc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/primefactorcalculator.PrimeFactorService/PrimeFactorCalc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrimeFactorServiceServer).PrimeFactorCalc(ctx, req.(*NumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrimeFactorService_ServiceDesc is the grpc.ServiceDesc for PrimeFactorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrimeFactorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "primefactorcalculator.PrimeFactorService",
	HandlerType: (*PrimeFactorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrimeFactorCalc",
			Handler:    _PrimeFactorService_PrimeFactorCalc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pf_calculator.proto",
}
