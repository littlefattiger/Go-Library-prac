package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"features/authentication"
	"features/data"
	ecpb "features/proto/echo"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callUnaryEcho(client ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("client.UnaryEcho(_) = _, %v: ", err)
	}
	fmt.Println("UnaryEcho: ", resp.Message)
}

func main() {
	flag.Parse()

	// 构建一个 PerRPCCredentials。
	// perRPC := oauth.NewOauthAccess(fetchToken())
	myAuth := authentication.NewMyAuth()
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca.crt"), "www.xiaomingdou.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(perRPC))
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(myAuth))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := ecpb.NewEchoClient(conn)

	callUnaryEcho(client, "hello world")
}

// fetchToken 获取授权信息
func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
