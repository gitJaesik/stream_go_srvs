package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	// pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1/stream_auth_serverv1"
	// pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1/stream_auth_serverv1"
	pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_auth_server/v1"
)

// pbsas "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_common/v1/stream_auth_serverv1"
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

	// creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.perm"), "x.test.example.com")
	// if err != nil {
	// 	log.Fatalf("failed to load credentials: %v", err)
	// }

	// opts := []grpc.DialOption{
	// 	grpc.WithTransportCredentials(creds),
	// }
	// opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
