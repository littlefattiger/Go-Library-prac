package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"

	"log"
	"os"
	pb "pro01/pb"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("did not connect: ", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("can not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}
