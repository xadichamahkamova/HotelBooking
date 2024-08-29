package bookingservice

import (
	pb "api-gateway/genproto/bookingpb"
	config "api-gateway/internal/pkg/load"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithBookingService(cfg config.Config) (*pb.BookingServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.BookingService.Host, cfg.BookingService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	bookingServiceClient := pb.NewBookingServiceClient(conn)
	return &bookingServiceClient, nil
}