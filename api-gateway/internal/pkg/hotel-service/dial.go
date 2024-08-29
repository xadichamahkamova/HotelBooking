package hotelservice

import (
	pb "api-gateway/genproto/hotelpb"
	config "api-gateway/internal/pkg/load"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithHotelService(cfg config.Config) (*pb.HotelServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.HotelService.Host, cfg.HotelService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	hotelServiceClient := pb.NewHotelServiceClient(conn)
	return &hotelServiceClient, nil
}