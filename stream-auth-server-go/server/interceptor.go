package server

import (
	"context"
	"errors"
	"sync"

	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
	"google.golang.org/grpc"
)

type ServiceGRPCExtension interface {
	GRPCExtension(ctx context.Context, fullMethodName string) (context.Context, error)
}

type StreamServerInterceptor struct {
	// redis
	// mongoDb sgldb.StreamGoLibDB
}

var once sync.Once
var streamServerInterceptor *StreamServerInterceptor

// NewStreamServerInterceptor ...
func NewStreamServerInterceptor() *StreamServerInterceptor {
	once.Do(func() {
		logger.Logger.Infow("NewStreamServerInterceptor", "sync.once.Do", "singleton")
		streamServerInterceptor = &StreamServerInterceptor{
			// redis? mongodb
		}
	})

	if streamServerInterceptor == nil {
		panic(errors.New("NewStreamServerInterceptor"))
	}

	return streamServerInterceptor
}

func (ssi *StreamServerInterceptor) AuthUnaryInterceptorHandler(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	switch info.FullMethod {
	case "/":
		logger.Logger.Infow("AuthUnaryInterceptorHandler", "", "")
		return handler(ctx, req)
	default:
		logger.Logger.Errorw("AuthUnaryInterceptorHandler", "switch", "default")
		return handler(ctx, req)
	}
	return nil, nil
}
