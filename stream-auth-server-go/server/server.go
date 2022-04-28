package server

import (
	"crypto/tls"

	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/config"
	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/service"
	sglconfig "github.com/gitJaesik/stream_go_srvs/streamgolib/config"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/data"
	sgldb "github.com/gitJaesik/stream_go_srvs/streamgolib/db"
	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
	"github.com/google/wire"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var Set = wire.NewSet(
	InitializeGrpcServer,
)

// InitializeGrpcServer ...
func InitializeGrpcServer(sglConfig *config.SGLConfig, mongoDb sgldb.StreamGoLibDB) *grpc.Server {

	cert, err := tls.LoadX509KeyPair(data.Path(sglconfig.GetCertFile()), data.Path(sglconfig.GetKeyFile()))
	if err != nil {
		logger.Logger.Errorw("InitializeGrpcServer", "failed to load key pair", err)
	}
	logger.Logger.Infow("InitializeGrpcServer", "sgldata.Path(sglconfig.GetCertFile())", data.Path(sglconfig.GetCertFile()))
	logger.Logger.Infow("InitializeGrpcServer", "sgldata.Path(sglconfig.GetKeyFile())", data.Path(sglconfig.GetKeyFile()))

	streamServerInterceptor := NewStreamServerInterceptor()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(streamServerInterceptor.AuthUnaryInterceptorHandler),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	// streamServerInterceptor := NewStreamServerInterceptor()

	// opts := []grpc.ServerOption{
	// 	grpc.UnaryInterceptor(streamServerInterceptor.AuthUnaryInterceptorHandler),
	// }

	// Create a gRPC server object
	grpcServer := grpc.NewServer(opts...)

	// Attach the ... service to the server
	pbsas.RegisterStreamAuthServiceServer(grpcServer, service.NewStreamAuthServerServer(mongoDb))

	return grpcServer
}
