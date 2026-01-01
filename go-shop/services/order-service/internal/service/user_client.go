package service

import (
	"context"
	"time"

	userpb "go-shop/order-service/proto"
	"google.golang.org/grpc"
)

// UserClient 封装 User Service 的 gRPC Client
type UserClient struct {
	client userpb.UserServiceClient
}

// NewUserClient 创建 User Service 客户端
func NewUserClient(addr string) *UserClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(), // 本地先用明文
	)
	if err != nil {
		panic(err)
	}

	return &UserClient{
		client: userpb.NewUserServiceClient(conn),
	}
}

// GetUser 调用 User Service
func (c *UserClient) GetUser(ctx context.Context, userID int64) (*userpb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return c.client.GetUser(ctx, &userpb.GetUserRequest{
		Id: userID,
	})
}
