package server

import (
	"context"
	"crypto/tls"
	"log"

	_ "github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/test/testdata"
	_ "github.com/gitJaesik/stream_go_srvs/streamgolib/config" // for save
	_ "github.com/gitJaesik/stream_go_srvs/streamgolib/data"   // for save
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/config"
	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/service"
	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/test/testdata"
	sglconfig "github.com/gitJaesik/stream_go_srvs/streamgolib/config"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/data"
	sgldb "github.com/gitJaesik/stream_go_srvs/streamgolib/db"
	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
	"github.com/google/wire"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var Set = wire.NewSet(
	InitializeGrpcServer,
)

// InitializeGrpcServer ...
func InitializeGrpcServer(sglConfig *config.SGLConfig, mongoDb sgldb.StreamGoLibDB) *grpc.Server {

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

// InitializeSecureGrpcServer ...
func InitializeSecureGrpcServer(sglConfig *config.SGLConfig, mongoDb sgldb.StreamGoLibDB) *grpc.Server {

	cert, err := tls.LoadX509KeyPair(data.Path(sglconfig.GetCertFile()), data.Path(sglconfig.GetKeyFile()))
	if err != nil {
		logger.Logger.Errorw("InitializeGrpcServer", "failed to load key pair", err)
		panic(err)
	}
	logger.Logger.Infow("InitializeGrpcServer", "sgldata.Path(sglconfig.GetCertFile())", data.Path(sglconfig.GetCertFile()))
	logger.Logger.Infow("InitializeGrpcServer", "sgldata.Path(sglconfig.GetKeyFile())", data.Path(sglconfig.GetKeyFile()))

	streamServerInterceptor := NewStreamServerInterceptor()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(streamServerInterceptor.AuthUnaryInterceptorHandler),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	// Create a gRPC server object
	grpcServer := grpc.NewServer(opts...)

	// Attach the ... service to the server
	pbsas.RegisterStreamAuthServiceServer(grpcServer, service.NewStreamAuthServerServer(mongoDb))

	return grpcServer
}

// InitializeGatewayMux ...
func InitializeGatewayMux() *runtime.ServeMux {

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8280", opts...)
	if err != nil {
		logger.Logger.Errorw("Failed to dial server", "failed to load key pair", err)
		panic(err)
	}

	err = pbsas.RegisterStreamAuthServiceHandler(context.Background(), mux, conn)
	if err != nil {
		panic(err)
	}

	return mux
}

// InitializeSecureGatewayMux ...
func InitializeSecureGatewayMux() *runtime.ServeMux {

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	creds, err := credentials.NewClientTLSFromFile(testdata.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithBlock(), grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8280", opts...)
	if err != nil {
		logger.Logger.Errorw("Failed to dial server", "failed to load key pair", err)
		panic(err)
	}

	err = pbsas.RegisterStreamAuthServiceHandler(context.Background(), mux, conn)
	if err != nil {
		panic(err)
	}

	return mux
}
