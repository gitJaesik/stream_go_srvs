package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/test/testdata"
	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
)

const (
	address = "localhost:8280"
)

func callUnaryEcho(client pbsas.StreamAuthServiceClient, er *pbsas.EchoRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := client.Echo(ctx, er)
	if err != nil {
		fmt.Printf("client.Echo()) = _, %v: ", err)
	}
	fmt.Println("Unary Echo : ", resp)
}

func main() {

	logger.Logger.Infow("test/grpc_unary_echo", "data.Path for cert", testdata.Path("x509/ca_cert.pem"))

	creds, err := credentials.NewClientTLSFromFile(testdata.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(address, opts...)
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbsas.NewStreamAuthServiceClient(conn)
	echoRequest := &pbsas.EchoRequest{
		Message: "echo test stream auth",
	}
	callUnaryEcho(c, echoRequest)

}
