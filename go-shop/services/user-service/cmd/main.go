package main

import (
	"log"
	"net"
	"net/http"

	"go-shop/user-service/internal/service"
	userpb "go-shop/user-service/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// ---------- gRPC ----------
	grpcServer := grpc.NewServer()
	userSvc := service.NewUserService()

	userpb.RegisterUserServiceServer(grpcServer, userSvc)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Println("gRPC server listening on :9000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}
	}()

	// ---------- HTTP ----------
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("HTTP server listening on :8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
