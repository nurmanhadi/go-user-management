package handler

import (
	"context"
	"user-management/internal/protobuf/pb"
	"user-management/internal/service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
func (h *UserHandler) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	result, err := h.userService.UserGetById(req.UserId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
