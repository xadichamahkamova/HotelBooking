package service

import (
	"context"
	pb "hotel-service/genproto/hotelpb"
	"hotel-service/internal/repository"
)

type HotelService struct {
	pb.UnimplementedHotelServiceServer
	Repo repository.IHotelRepository
}

func NewHotelService(repo repository.IHotelRepository) *HotelService {
	return &HotelService{
		Repo: repo,
	}
}

func (s *HotelService) ListOfHotel(ctx context.Context, req *pb.ListOfHotelRequest) (*pb.ListOfHotelResponse, error) {
	return s.Repo.ListOfHotel()
}

func (s *HotelService) GetDetailsOfHotel(ctx context.Context, req *pb.GetDetailsOfHotelRequest) (*pb.GetDetailsOfHotelResponse, error) {
	return s.Repo.GetDetailsOfHotel(req)
}

func (s *HotelService) GetAvailabilityRooms(ctx context.Context, req *pb.GetAvailabilityRoomsRequest) (*pb.GetAvailabilityRoomsResponse, error) {
	return s.Repo.GetAvailabilityRooms(req)
}