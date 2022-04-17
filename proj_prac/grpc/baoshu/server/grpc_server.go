package main

import (
	"baoshu/service"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 1. new一个grpc的server
	tls, err := credentials.NewServerTLSFromFile("./keys/server.crt", "./keys/server_no_password.key")
	if err != nil {
		log.Fatal("服务端获取证书失败: ", err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(tls))
	// rpcServer := grpc.NewServer()
	// 2. 将刚刚我们新建的ProdService注册进去
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))

	// 3. 新建一个listener，以tcp方式监听8082端口
	// listener, err := net.Listen("tcp", ":8082")
	// if err != nil {
	// 	log.Fatal("服务监听端口失败", err)
	// }

	// 4. 运行rpcServer，传入listener
	// _ = rpcServer.Serve(listener)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		rpcServer.ServeHTTP(writer, request)
	})

	// 5. 定义httpServer，监听8082
	httpServer := http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	// 6. 以https形式监听httpServer
	// httpServer.ListenAndServeTLS("grpc_server/keys/server.crt", "grpc_server/keys/server_no_password.key")
	httpServer.ListenAndServe()
}
