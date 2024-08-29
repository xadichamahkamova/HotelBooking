package registerservice

import (
	"fmt"
	config "hotel-service/internal/pkg/load"
	"hotel-service/internal/service"
	"hotel-service/logger"
	"net"

	pb "hotel-service/genproto/hotelpb"

	"google.golang.org/grpc"
)

type Service struct {
	Service *service.HotelService
}

func NewGRPCHotelService(service *service.HotelService) *Service {
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
	pb.RegisterHotelServiceServer(grpcServer, s.Service)
	logger.Info("Hotel-service listening on :", cfg.Service.Port)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func GracefufShutdown() {
	grpcServer.GracefulStop()
}