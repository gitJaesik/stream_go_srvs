// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: stream_auth_server/v1/stream_auth_server.proto

package stream_auth_serverv1

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

// StreamAuthServiceClient is the client API for StreamAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamAuthServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	// rpc Echo(EchoRequest) returns (EchoResponse) {
	//   option (google.api.http) = {
	//     post: "/v1/example/echo"
	//     body: "*"
	//   };
	// }
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	CreatePlayerInfo(ctx context.Context, in *CreatePlayerInfoRequest, opts ...grpc.CallOption) (*CreatePlayerInfoResponse, error)
	GetPlayerInfo(ctx context.Context, in *GetPlayerInfoRequest, opts ...grpc.CallOption) (*GetPlayerInfoResponse, error)
	UpdatePlayer(ctx context.Context, in *GetPlayerInfoRequest, opts ...grpc.CallOption) (*GetPlayerInfoResponse, error)
}

type streamAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamAuthServiceClient(cc grpc.ClientConnInterface) StreamAuthServiceClient {
	return &streamAuthServiceClient{cc}
}

func (c *streamAuthServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/stream_auth_server.v1.StreamAuthService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamAuthServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/stream_auth_server.v1.StreamAuthService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamAuthServiceClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, "/stream_auth_server.v1.StreamAuthService/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamAuthServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/stream_auth_server.v1.StreamAuthService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamAuthServiceClient) CreatePlayerInfo(ctx context.Context, in *CreatePlayerInfoRequest, opts ...grpc.CallOption) (*CreatePlayerInfoResponse, error) {
	out := new(CreatePlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/stream_auth_server.v1.StreamAuthService/CreatePlayerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamAuthServiceClient) GetPlayerInfo(ctx context.Context, in *GetPlayerInfoRequest, opts ...grpc.CallOption) (*GetPlayerInfoResponse, error) {
	out := new(GetPlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/stream_auth_server.v1.StreamAuthService/GetPlayerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamAuthServiceClient) UpdatePlayer(ctx context.Context, in *GetPlayerInfoRequest, opts ...grpc.CallOption) (*GetPlayerInfoResponse, error) {
	out := new(GetPlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/stream_auth_server.v1.StreamAuthService/UpdatePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamAuthServiceServer is the server API for StreamAuthService service.
// All implementations should embed UnimplementedStreamAuthServiceServer
// for forward compatibility
type StreamAuthServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	// rpc Echo(EchoRequest) returns (EchoResponse) {
	//   option (google.api.http) = {
	//     post: "/v1/example/echo"
	//     body: "*"
	//   };
	// }
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	CreatePlayerInfo(context.Context, *CreatePlayerInfoRequest) (*CreatePlayerInfoResponse, error)
	GetPlayerInfo(context.Context, *GetPlayerInfoRequest) (*GetPlayerInfoResponse, error)
	UpdatePlayer(context.Context, *GetPlayerInfoRequest) (*GetPlayerInfoResponse, error)
}

// UnimplementedStreamAuthServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStreamAuthServiceServer struct {
}

func (UnimplementedStreamAuthServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedStreamAuthServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedStreamAuthServiceServer) SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedStreamAuthServiceServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedStreamAuthServiceServer) CreatePlayerInfo(context.Context, *CreatePlayerInfoRequest) (*CreatePlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayerInfo not implemented")
}
func (UnimplementedStreamAuthServiceServer) GetPlayerInfo(context.Context, *GetPlayerInfoRequest) (*GetPlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerInfo not implemented")
}
func (UnimplementedStreamAuthServiceServer) UpdatePlayer(context.Context, *GetPlayerInfoRequest) (*GetPlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlayer not implemented")
}

// UnsafeStreamAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamAuthServiceServer will
// result in compilation errors.
type UnsafeStreamAuthServiceServer interface {
	mustEmbedUnimplementedStreamAuthServiceServer()
}

func RegisterStreamAuthServiceServer(s grpc.ServiceRegistrar, srv StreamAuthServiceServer) {
	s.RegisterService(&StreamAuthService_ServiceDesc, srv)
}

func _StreamAuthService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamAuthServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_auth_server.v1.StreamAuthService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAuthServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamAuthService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamAuthServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_auth_server.v1.StreamAuthService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAuthServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamAuthService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamAuthServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_auth_server.v1.StreamAuthService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAuthServiceServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamAuthService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamAuthServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_auth_server.v1.StreamAuthService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAuthServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamAuthService_CreatePlayerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamAuthServiceServer).CreatePlayerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_auth_server.v1.StreamAuthService/CreatePlayerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAuthServiceServer).CreatePlayerInfo(ctx, req.(*CreatePlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamAuthService_GetPlayerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamAuthServiceServer).GetPlayerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_auth_server.v1.StreamAuthService/GetPlayerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAuthServiceServer).GetPlayerInfo(ctx, req.(*GetPlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamAuthService_UpdatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamAuthServiceServer).UpdatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream_auth_server.v1.StreamAuthService/UpdatePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAuthServiceServer).UpdatePlayer(ctx, req.(*GetPlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamAuthService_ServiceDesc is the grpc.ServiceDesc for StreamAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream_auth_server.v1.StreamAuthService",
	HandlerType: (*StreamAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _StreamAuthService_Echo_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _StreamAuthService_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _StreamAuthService_SignOut_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _StreamAuthService_SignUp_Handler,
		},
		{
			MethodName: "CreatePlayerInfo",
			Handler:    _StreamAuthService_CreatePlayerInfo_Handler,
		},
		{
			MethodName: "GetPlayerInfo",
			Handler:    _StreamAuthService_GetPlayerInfo_Handler,
		},
		{
			MethodName: "UpdatePlayer",
			Handler:    _StreamAuthService_UpdatePlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stream_auth_server/v1/stream_auth_server.proto",
}
