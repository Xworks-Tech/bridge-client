// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bridge

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

// KafkaStreamClient is the client API for KafkaStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaStreamClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (KafkaStream_SubscribeClient, error)
}

type kafkaStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaStreamClient(cc grpc.ClientConnInterface) KafkaStreamClient {
	return &kafkaStreamClient{cc}
}

func (c *kafkaStreamClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (KafkaStream_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &KafkaStream_ServiceDesc.Streams[0], "/bridge.KafkaStream/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &kafkaStreamSubscribeClient{stream}
	return x, nil
}

type KafkaStream_SubscribeClient interface {
	Send(*PublishRequest) error
	Recv() (*KafkaResponse, error)
	grpc.ClientStream
}

type kafkaStreamSubscribeClient struct {
	grpc.ClientStream
}

func (x *kafkaStreamSubscribeClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kafkaStreamSubscribeClient) Recv() (*KafkaResponse, error) {
	m := new(KafkaResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KafkaStreamServer is the server API for KafkaStream service.
// All implementations must embed UnimplementedKafkaStreamServer
// for forward compatibility
type KafkaStreamServer interface {
	Subscribe(KafkaStream_SubscribeServer) error
	mustEmbedUnimplementedKafkaStreamServer()
}

// UnimplementedKafkaStreamServer must be embedded to have forward compatible implementations.
type UnimplementedKafkaStreamServer struct {
}

func (UnimplementedKafkaStreamServer) Subscribe(KafkaStream_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedKafkaStreamServer) mustEmbedUnimplementedKafkaStreamServer() {}

// UnsafeKafkaStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaStreamServer will
// result in compilation errors.
type UnsafeKafkaStreamServer interface {
	mustEmbedUnimplementedKafkaStreamServer()
}

func RegisterKafkaStreamServer(s grpc.ServiceRegistrar, srv KafkaStreamServer) {
	s.RegisterService(&KafkaStream_ServiceDesc, srv)
}

func _KafkaStream_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KafkaStreamServer).Subscribe(&kafkaStreamSubscribeServer{stream})
}

type KafkaStream_SubscribeServer interface {
	Send(*KafkaResponse) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type kafkaStreamSubscribeServer struct {
	grpc.ServerStream
}

func (x *kafkaStreamSubscribeServer) Send(m *KafkaResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kafkaStreamSubscribeServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KafkaStream_ServiceDesc is the grpc.ServiceDesc for KafkaStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bridge.KafkaStream",
	HandlerType: (*KafkaStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _KafkaStream_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/kafka.proto",
}
