package service

import (
	"context"

	pb "github.com/Parthiba-Hazra/user-management-microservice/api"
)

type UserService struct {
	// Add dependencies and database connections.
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {

}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

}
