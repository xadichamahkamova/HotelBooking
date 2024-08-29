package repository

import (
	pb "hotel-service/genproto/hotelpb"
)

type IHotelRepository interface {
	ListOfHotel() (*pb.ListOfHotelResponse,  error)
	GetDetailsOfHotel(req *pb.GetDetailsOfHotelRequest) (*pb.GetDetailsOfHotelResponse, error)
	GetAvailabilityRooms(req *pb.GetAvailabilityRoomsRequest) (*pb.GetAvailabilityRoomsResponse, error)
}