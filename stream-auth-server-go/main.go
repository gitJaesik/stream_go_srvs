package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/gitJaesik/stream_go_srvs/streamgolib"
	sglmongo "github.com/gitJaesik/stream_go_srvs/streamgolib/db/mongodb"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8280", "gRPC server endpoint")
)

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

	if err := srv.Server(lis); err != nil {
		logger.Logger.Errorw("main", "server run error", err)
	}

}
