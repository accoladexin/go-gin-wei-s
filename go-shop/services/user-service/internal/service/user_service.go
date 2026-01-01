package service

import (
	"context"
	userpb "go-shop/user-service/proto"
	"log"
)

type UserService struct {
	userpb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(
	ctx context.Context,
	req *userpb.GetUserRequest,
) (*userpb.GetUserResponse, error) {

	log.Println("GetUser: ", req.Id)
	// 先写死，后面接数据库
	return &userpb.GetUserResponse{
		Id:    req.Id,
		Name:  "Alice",
		Email: "alice@example.com",
	}, nil
}
