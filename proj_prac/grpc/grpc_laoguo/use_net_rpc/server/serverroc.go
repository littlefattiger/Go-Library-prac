package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
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
	http.Serve(lis, nil)

}
