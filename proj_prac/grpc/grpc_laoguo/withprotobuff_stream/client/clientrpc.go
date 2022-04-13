package main

import (
	"context"

	"google.golang.org/grpc/credentials/insecure"

	"log"
	pb "pro01/pb"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/grpclb"
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

	reqstreamData := &pb.StreamReqData{Data: "aaa"}
	res, _ := c.GetStream(context.Background(), reqstreamData)
	for {
		aa, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(aa)
	}

}
