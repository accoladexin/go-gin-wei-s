package service

import (
	"context"
	"log"
	"time"

	orderpb "go-shop/order-service/proto"
)

// OrderService 是 Order 微服务的核心实现
type OrderService struct {
	orderpb.UnimplementedOrderServiceServer

	userClient    *UserClient
	productClient *ProductClient
}

// NewOrderService 构造函数（依赖注入）
func NewOrderService(
	userClient *UserClient,
	productClient *ProductClient,
) *OrderService {
	return &OrderService{
		userClient:    userClient,
		productClient: productClient,
	}
}

// CreateOrder 创建订单（核心流程）
func (s *OrderService) CreateOrder(
	ctx context.Context,
	req *orderpb.CreateOrderRequest,
) (*orderpb.CreateOrderResponse, error) {

	// 1️⃣ 调用 User Service
	user, err := s.userClient.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	// 2️⃣ 调用 Product Service
	product, err := s.productClient.GetProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	// 3️⃣ 模拟生成订单 ID（真实场景是 DB 自增）
	orderID := time.Now().Unix()

	log.Printf(
		"create order success: user=%s product=%s",
		user.Name,
		product.Name,
	)

	// 4️⃣ 返回结果（注意：冗余字段）
	return &orderpb.CreateOrderResponse{
		OrderId:     orderID,
		UserName:    user.Name,
		ProductName: product.Name,
		Price:       product.Price,
	}, nil
}
