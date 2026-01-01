package service

import (
	"context"
	productpb "go-shop/product-service/proto"
	"log"
)

type ProductService struct {
	productpb.UnimplementedProductServiceServer
	userClient *UserClient
}

func NewProductService(userClient *UserClient) *ProductService {
	return &ProductService{
		userClient: userClient,
	}
}

func (s *ProductService) GetProduct(
	ctx context.Context,
	req *productpb.GetProductRequest,
) (*productpb.GetProductResponse, error) {
	user, _ := s.userClient.GetUser(ctx, 1)
	log.Println("get user from user-service:", user.Name)

	return &productpb.GetProductResponse{
		Id:    req.Id,
		Name:  "MacBook Pro",
		Price: 19999,
	}, nil
}
