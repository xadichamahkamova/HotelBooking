package service

import (
	"context"
	pb "user-service/genproto/userpb"
	"user-service/internal/repository"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return s.Repo.RegisterUser(req)
}

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return s.Repo.LoginUser(req)
}