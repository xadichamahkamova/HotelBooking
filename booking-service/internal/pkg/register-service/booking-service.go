package registerservice

import (
	config "booking-service/internal/pkg/load"
	"booking-service/internal/service"
	"booking-service/logger"
	"fmt"
	"net"

	pb "booking-service/genproto/bookingpb"

	"google.golang.org/grpc"
)

type Service struct {
	Service *service.BookingService
}

func NewGRPCHotelService(service *service.BookingService) *Service {
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
	pb.RegisterBookingServiceServer(grpcServer, s.Service)
	logger.Info("Booking-service listening on :", cfg.Service.Port)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func GracefufShutdown() {
	grpcServer.GracefulStop()
}
