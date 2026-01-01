package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-shop/order-service/internal/handler"
	"go-shop/order-service/internal/service"
	orderpb "go-shop/order-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// ========== gRPC Server ==========
	grpcServer := grpc.NewServer()

	userClient := service.NewUserClient("user-service:9000")
	productClient := service.NewProductClient("product-service:9001")

	orderSvc := service.NewOrderService(userClient, productClient)
	orderpb.RegisterOrderServiceServer(grpcServer, orderSvc)

	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("Order gRPC server listening on :9002")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	// ========== HTTP Server ==========
	r := gin.Default()

	orderHandler := handler.NewOrderHTTPHandler("localhost:9002")
	r.POST("/orders", orderHandler.CreateOrder)

	log.Println("Order HTTP server listening on :8002")
	if err := http.ListenAndServe(":8002", r); err != nil {
		log.Fatal(err)
	}
}
