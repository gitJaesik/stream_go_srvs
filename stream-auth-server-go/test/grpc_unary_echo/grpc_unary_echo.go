package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
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

	opts := []grpc.DialOption{
		grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(address, opts...)
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
