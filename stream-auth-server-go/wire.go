//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/config"
	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/server"
	sgldb "github.com/gitJaesik/stream_go_srvs/streamgolib/db"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// initializeConfig ...
func initializeConfig() *config.SGLConfig {
	wire.Build(
		config.Set,
	)
	return nil
}

func initializeServer(
	sglConfig *config.SGLConfig, mongodb sgldb.StreamGoLibDB) *grpc.Server {
	wire.Build(
		server.Set,
	)
	return nil
}

func initializeGatewayMux() *runtime.ServeMux {
	wire.Build(
		server.GatewaySet,
	)
	return nil
}
