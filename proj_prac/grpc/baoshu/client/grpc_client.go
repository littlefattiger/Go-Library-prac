package main

import (
	"baoshu/client/service"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 1. 新建连接，端口是服务端开放的8082端口
	// 并且添加grpc.WithInsecure()，不然没有证书会报错
	tls, err := credentials.NewClientTLSFromFile("./keys/server.crt", "testRoot")

	if err != nil {
		log.Fatal("客户端获取证书失败: ", err)
	}

	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(tls))
	// conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	// 退出时关闭链接
	defer conn.Close()

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	productServiceClient := service.NewProdServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	resp, err := productServiceClient.GetProductStock(context.Background(), &service.ProductRequest{ProdId: 233})
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	fmt.Println("调用gRPC方法成功，ProdStock = ", resp.ProdStock)
}
