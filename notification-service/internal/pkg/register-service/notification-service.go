package registerservice


import (
	config "notification-service/internal/pkg/load"
	"notification-service/internal/service"
	"notification-service/logger"
	"fmt"
	"net"

	pb "notification-service/genproto/notificationpb"

	"google.golang.org/grpc"
)

type Service struct {
	Service *service.NotificationService
}

func NewGRPCNotificationService(service *service.NotificationService) *Service {
	return &Service{
		Service: service,
	}
}

var grpcServer = grpc.NewServer()

func (s *Service) RUN(cfg config.Config) error {

	addr := fmt.Sprintf(":%d", cfg.Service.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	pb.RegisterNotificationServiceServer(grpcServer, s.Service)
	logger.Info("Notification-service listening on :", cfg.Service.Port)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func GracefufShutdown() {
	grpcServer.GracefulStop()
}
