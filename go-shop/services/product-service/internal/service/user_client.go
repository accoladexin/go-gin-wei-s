package service

import (
	"context"
	"log"
	"time"

	userpb "go-shop/product-service/proto"
	"google.golang.org/grpc"
)

type UserClient struct {
	client userpb.UserServiceClient
}

// 主要grpc的地址
func NewUserClient(addr string) *UserClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &UserClient{
		client: userpb.NewUserServiceClient(conn),
	}
}

// 方法
func (c *UserClient) GetUser(ctx context.Context, userID int64) (*userpb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return c.client.GetUser(ctx, &userpb.GetUserRequest{
		Id: userID,
	})
}
