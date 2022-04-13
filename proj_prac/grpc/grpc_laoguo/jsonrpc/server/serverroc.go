package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Arith struct {
}

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func (arith *Arith) Multiply(req ArithRequest, res *ArithResponse) error {

	res.Pro = req.A * req.B

	return nil
}

func (arith *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}

	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	rpc.Register(new(Arith))
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection")
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in comming")
			jsonrpc.ServeConn(conn)

		}(conn)
	}

}
