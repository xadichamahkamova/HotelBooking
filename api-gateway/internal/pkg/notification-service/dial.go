package notificationservice

import (
	pb "api-gateway/genproto/notificationpb"
	config "api-gateway/internal/pkg/load"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithNotificationService(cfg config.Config) (*pb.NotificationServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.NotificationService.Host, cfg.NotificationService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	notificationServiceClient := pb.NewNotificationServiceClient(conn)
	return &notificationServiceClient, nil
}