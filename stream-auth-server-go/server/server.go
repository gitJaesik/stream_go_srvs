package server

import (
	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/config"
	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/service"
	sgldb "github.com/gitJaesik/stream_go_srvs/streamgolib/db"
	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
	"github.com/google/wire"

	"google.golang.org/grpc"
)

var Set = wire.NewSet(
	InitializeGrpcServer,
)

// InitializeGrpcServer ...
func InitializeGrpcServer(sglConfig *config.SGLConfig, mongoDb sgldb.StreamGoLibDB) *grpc.Server {

	/*
		cert, err := tls.LoadX509KeyPair(sgldata.Path(sglconfig.GetCertFile()), sgldata.Path(sglconfig.GetKeyFile()))
		if err != nil {
			logger.Logger.Errorw("InitializeGrpcServer", "failed to load key pair", err)
		}
		logger.Logger.Infow("InitializeGrpcServer", "sgldata.Path(sglconfig.GetCertFile())", sgldata.Path(sglconfig.GetCertFile()))
		logger.Logger.Infow("InitializeGrpcServer", "sgldata.Path(sglconfig.GetKeyFile())", sgldata.Path(sglconfig.GetKeyFile()))

		streamServerInterceptor := NewStreamServerInterceptor()

		opts := []grpc.ServerOption{
			grpc.UnaryInterceptor(streamServerInterceptor.AuthUnaryInterceptorHandler),
			grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		}
	*/

	streamServerInterceptor := NewStreamServerInterceptor()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(streamServerInterceptor.AuthUnaryInterceptorHandler),
	}

	// Create a gRPC server object
	grpcServer := grpc.NewServer(opts...)

	// Attach the ... service to the server
	pbsas.RegisterStreamAuthServiceServer(grpcServer, service.NewStreamAuthServerServer(mongoDb))

	return grpcServer
}
