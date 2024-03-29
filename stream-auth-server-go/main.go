package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/server"
	"github.com/gitJaesik/stream_go_srvs/streamgolib" // for save
	sglmongo "github.com/gitJaesik/stream_go_srvs/streamgolib/db/mongodb"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
)

// var (
// 	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8280", "gRPC server endpoint")
// )

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

	go func() {
		httpSwagger := server.GetOpenAPIHandler()
		http.ListenAndServe(":8282", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpSwagger.ServeHTTP(w, r)
		}))
	}()

	// if err := run(); err != nil {
	// 	glog.Fatal(err)
	// }
	gwMux := initializeGatewayMux()
	http.ListenAndServe(":"+sglConfig.GetAuthServerGwPort(), gwMux)

}
