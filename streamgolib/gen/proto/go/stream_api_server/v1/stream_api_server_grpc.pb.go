// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: stream_api_server/v1/stream_api_server.proto

package stream_api_serverv1

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

// StreamApiServiceClient is the client API for StreamApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamApiServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	SpeechNoteStream(ctx context.Context, opts ...grpc.CallOption) (StreamApiService_SpeechNoteStreamClient, error)
	CreateNoteWithSpeechNote(ctx context.Context, in *CreateNoteWithSpeechNoteRequest, opts ...grpc.CallOption) (*CreateNoteWithSpeechNoteResponse, error)
}

type streamApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamApiServiceClient(cc grpc.ClientConnInterface) StreamApiServiceClient {
	return &streamApiServiceClient{cc}
}

func (c *streamApiServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/stream_api_server.v1.StreamApiService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamApiServiceClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, "/stream_api_server.v1.StreamApiService/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamApiServiceClient) SpeechNoteStream(ctx context.Context, opts ...grpc.CallOption) (StreamApiService_SpeechNoteStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamApiService_ServiceDesc.Streams[0], "/stream_api_server.v1.StreamApiService/SpeechNoteStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamApiServiceSpeechNoteStreamClient{stream}
	return x, nil
}

type StreamApiService_SpeechNoteStreamClient interface {
	Send(*SpeechNoteStreamRequest) error
	CloseAndRecv() (*SpeechNoteStreamResponse, error)
	grpc.ClientStream
}

type streamApiServiceSpeechNoteStreamClient struct {
	grpc.ClientStream
}

func (x *streamApiServiceSpeechNoteStreamClient) Send(m *SpeechNoteStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamApiServiceSpeechNoteStreamClient) CloseAndRecv() (*SpeechNoteStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SpeechNoteStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamApiServiceClient) CreateNoteWithSpeechNote(ctx context.Context, in *CreateNoteWithSpeechNoteRequest, opts ...grpc.CallOption) (*CreateNoteWithSpeechNoteResponse, error) {
	out := new(CreateNoteWithSpeechNoteResponse)
	err := c.cc.Invoke(ctx, "/stream_api_server.v1.StreamApiService/CreateNoteWithSpeechNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamApiServiceServer is the server API for StreamApiService service.
// All implementations should embed UnimplementedStreamApiServiceServer
// for forward compatibility
type StreamApiServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	SpeechNoteStream(StreamApiService_SpeechNoteStreamServer) error
	CreateNoteWithSpeechNote(context.Context, *CreateNoteWithSpeechNoteRequest) (*CreateNoteWithSpeechNoteResponse, error)
}

// UnimplementedStreamApiServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStreamApiServiceServer struct {
}

func (UnimplementedStreamApiServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedStreamApiServiceServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedStreamApiServiceServer) SpeechNoteStream(StreamApiService_SpeechNoteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SpeechNoteStream not implemented")
}
func (UnimplementedStreamApiServiceServer) CreateNoteWithSpeechNote(context.Context, *CreateNoteWithSpeechNoteRequest) (*CreateNoteWithSpeechNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNoteWithSpeechNote not implemented")
}

// UnsafeStreamApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamApiServiceServer will
// result in compilation errors.
type UnsafeStreamApiServiceServer interface {
	mustEmbedUnimplementedStreamApiServiceServer()
}

func RegisterStreamApiServiceServer(s grpc.ServiceRegistrar, srv StreamApiServiceServer) {
	s.RegisterService(&StreamApiService_ServiceDesc, srv)
}

func _StreamApiService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamApiServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_api_server.v1.StreamApiService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamApiServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamApiService_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamApiServiceServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_api_server.v1.StreamApiService/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamApiServiceServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamApiService_SpeechNoteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamApiServiceServer).SpeechNoteStream(&streamApiServiceSpeechNoteStreamServer{stream})
}

type StreamApiService_SpeechNoteStreamServer interface {
	SendAndClose(*SpeechNoteStreamResponse) error
	Recv() (*SpeechNoteStreamRequest, error)
	grpc.ServerStream
}

type streamApiServiceSpeechNoteStreamServer struct {
	grpc.ServerStream
}

func (x *streamApiServiceSpeechNoteStreamServer) SendAndClose(m *SpeechNoteStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamApiServiceSpeechNoteStreamServer) Recv() (*SpeechNoteStreamRequest, error) {
	m := new(SpeechNoteStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamApiService_CreateNoteWithSpeechNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteWithSpeechNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamApiServiceServer).CreateNoteWithSpeechNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_api_server.v1.StreamApiService/CreateNoteWithSpeechNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamApiServiceServer).CreateNoteWithSpeechNote(ctx, req.(*CreateNoteWithSpeechNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamApiService_ServiceDesc is the grpc.ServiceDesc for StreamApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream_api_server.v1.StreamApiService",
	HandlerType: (*StreamApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _StreamApiService_Echo_Handler,
		},
		{
			MethodName: "CreateNote",
			Handler:    _StreamApiService_CreateNote_Handler,
		},
		{
			MethodName: "CreateNoteWithSpeechNote",
			Handler:    _StreamApiService_CreateNoteWithSpeechNote_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SpeechNoteStream",
			Handler:       _StreamApiService_SpeechNoteStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "stream_api_server/v1/stream_api_server.proto",
}