//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/config"
	sgldb "github.com/gitJaesik/stream_go_srvs/streamgolib/db"
	"github.com/google/wire"
)

// initializeConfig ...
func initializeConfig() *config.SGLConfig {
	wire.Build(
		config.Set,
	)
	return nil
}

func initializeServer(
	sglConfig *config.SGLConfig, sgldb.StreamGoLibDB) *grpc.Server {
	wire.Build(
		server.Set,
	)
	return nil
}
