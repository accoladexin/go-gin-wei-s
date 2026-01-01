package service

import (
	"context"
	productpb "go-shop/product-service/proto"
)

type ProductService struct {
	productpb.UnimplementedProductServiceServer
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetProduct(
	ctx context.Context,
	req *productpb.GetProductRequest,
) (*productpb.GetProductResponse, error) {

	return &productpb.GetProductResponse{
		Id:    req.Id,
		Name:  "MacBook Pro",
		Price: 19999,
	}, nil
}
