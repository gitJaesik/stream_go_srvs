package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/test/testdata"
	"github.com/gitJaesik/stream_go_srvs/streamgolib" // for save
	sglmongo "github.com/gitJaesik/stream_go_srvs/streamgolib/db/mongodb"
	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// var (
// 	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8280", "gRPC server endpoint")
// )

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err := pbsas.RegisterStreamAuthServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)

	// opts := []grpc.DialOption{
	// 	grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()),
	// }

	creds, err := credentials.NewClientTLSFromFile(testdata.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithBlock(), grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8280", opts...)
	if err != nil {
		fmt.Printf("Failed to dial server: %v", err)
		return err
	}

	err = pbsas.RegisterStreamAuthServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()

	streamgolib.SetProductName("stream-auth")
	fmt.Println("run " + streamgolib.ProductName())

	sglConfig := initializeConfig()
	logger.Logger.Infow("main", "stream go lib config", sglConfig)
	mongoDb := sglmongo.NewMongoDB()

	err := mongoDb.MongoPing(context.Background())
	if err != nil {
		logger.Logger.Errorw("main", "mongodb connection error", "server shutdown")
		return
	}
	logger.Logger.Infow("main", "sglConfig.GetAuthPort()", sglConfig.GetAuthServerGrpcPort())

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":"+sglConfig.GetAuthServerGrpcPort())
	if err != nil {
		logger.Logger.Errorw("main", "tcp listen error", time.Second)
		panic(err)
	}

	srv := initializeServer(sglConfig, mongoDb)

	// if err := srv.Serve(lis); err != nil {
	// 	logger.Logger.Errorw("main", "server run error", err)
	// }

	go func() {
		if err := srv.Serve(lis); err != nil {
			logger.Logger.Errorw("main", "server run error", err)
		}

	}()

	if err := run(); err != nil {
		glog.Fatal(err)
	}

}
