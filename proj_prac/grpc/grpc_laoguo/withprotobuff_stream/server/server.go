package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "pro01/pb"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) GetStream(req *pb.StreamReqData, res pb.Greeter_GetStreamServer) error {

	i := 0
	for {
		i++
		res.Send(&pb.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
