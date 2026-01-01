package main

import (
	"log"
	"net"

	"go-shop/product-service/internal/service"
	productpb "go-shop/product-service/proto"

	"google.golang.org/grpc"
)

func main() {
	// 1️⃣ 创建 gRPC Server
	grpcServer := grpc.NewServer()

	// 2️⃣ 创建 gRPC Client（依赖）
	//userClient := service.NewUserClient("localhost:9000")

	// 3️⃣ 创建 Product Service（注入依赖）
	productSvc := service.NewProductService()

	// 4️⃣ 注册服务
	productpb.RegisterProductServiceServer(grpcServer, productSvc)

	// 5️⃣ 监听端口
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Product gRPC server listening on :9001")

	// 6️⃣ 启动服务
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
