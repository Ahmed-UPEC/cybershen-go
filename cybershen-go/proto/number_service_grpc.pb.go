// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/number_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NumberService_BiDirectionalStreaming_FullMethodName = "/proto.NumberService/BiDirectionalStreaming"
)

// NumberServiceClient is the client API for NumberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Le reste du contenu .proto
type NumberServiceClient interface {
	BiDirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[NumberRequest, NumberResponse], error)
}

type numberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNumberServiceClient(cc grpc.ClientConnInterface) NumberServiceClient {
	return &numberServiceClient{cc}
}

func (c *numberServiceClient) BiDirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[NumberRequest, NumberResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NumberService_ServiceDesc.Streams[0], NumberService_BiDirectionalStreaming_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NumberRequest, NumberResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NumberService_BiDirectionalStreamingClient = grpc.BidiStreamingClient[NumberRequest, NumberResponse]

// NumberServiceServer is the server API for NumberService service.
// All implementations must embed UnimplementedNumberServiceServer
// for forward compatibility.
//
// Le reste du contenu .proto
type NumberServiceServer interface {
	BiDirectionalStreaming(grpc.BidiStreamingServer[NumberRequest, NumberResponse]) error
	mustEmbedUnimplementedNumberServiceServer()
}

// UnimplementedNumberServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNumberServiceServer struct{}

func (UnimplementedNumberServiceServer) BiDirectionalStreaming(grpc.BidiStreamingServer[NumberRequest, NumberResponse]) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionalStreaming not implemented")
}
func (UnimplementedNumberServiceServer) mustEmbedUnimplementedNumberServiceServer() {}
func (UnimplementedNumberServiceServer) testEmbeddedByValue()                       {}

// UnsafeNumberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NumberServiceServer will
// result in compilation errors.
type UnsafeNumberServiceServer interface {
	mustEmbedUnimplementedNumberServiceServer()
}

func RegisterNumberServiceServer(s grpc.ServiceRegistrar, srv NumberServiceServer) {
	// If the following call pancis, it indicates UnimplementedNumberServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NumberService_ServiceDesc, srv)
}

func _NumberService_BiDirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NumberServiceServer).BiDirectionalStreaming(&grpc.GenericServerStream[NumberRequest, NumberResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NumberService_BiDirectionalStreamingServer = grpc.BidiStreamingServer[NumberRequest, NumberResponse]

// NumberService_ServiceDesc is the grpc.ServiceDesc for NumberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NumberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NumberService",
	HandlerType: (*NumberServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BiDirectionalStreaming",
			Handler:       _NumberService_BiDirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/number_service.proto",
}
