package repository

import (
	pb "user-service/genproto/userpb"
)

type IUserRepository interface {
	RegisterUser(req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
	LoginUser(req *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
}