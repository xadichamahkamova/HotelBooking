package userservice

import (
	pb "api-gateway/genproto/userpb"
	config "api-gateway/internal/pkg/load"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithUserService(cfg config.Config) (*pb.UserServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.UserService.Host, cfg.UserService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewUserServiceClient(conn)
	return &userServiceClient, nil
}