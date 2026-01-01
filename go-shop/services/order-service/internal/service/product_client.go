package service

import (
	"context"
	"time"

	productpb "go-shop/order-service/proto"
	"google.golang.org/grpc"
)

// ProductClient 封装 Product Service 的 gRPC Client
type ProductClient struct {
	client productpb.ProductServiceClient
}

// NewProductClient 创建 Product Service 客户端
func NewProductClient(addr string) *ProductClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	return &ProductClient{
		client: productpb.NewProductServiceClient(conn),
	}
}

// GetProduct 调用 Product Service
func (c *ProductClient) GetProduct(ctx context.Context, productID int64) (*productpb.GetProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return c.client.GetProduct(ctx, &productpb.GetProductRequest{
		Id: productID,
	})
}
