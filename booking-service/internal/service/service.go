package service

import (
	pb "booking-service/genproto/bookingpb"
	"booking-service/internal/repository"
	"context"
)

type BookingService struct {
	pb.UnimplementedBookingServiceServer
	Repo repository.IBookingRepository
}

func NewBookingService(repo repository.IBookingRepository) *BookingService {
	return &BookingService{
		Repo: repo,
	}
}

func (s *BookingService) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.Booking, error) {
	return s.Repo.CreateBooking(req)
}

func (s *BookingService) GetDetailsOfBooking(ctx context.Context, req *pb.GetDetailsOfBookingRequest) (*pb.Booking, error) {
	return s.Repo.GetDetailsOfBooking(req)
}

func (s *BookingService) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.Booking, error) {
	return s.Repo.UpdateBooking(req)
}

func (s *BookingService) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error) {
	return s.Repo.CancelBooking(req)
}

func (s *BookingService) ListBookingOfUsers(ctx context.Context, req *pb.ListBookingOfUserRequest) (*pb.ListBookingOfUserResponse, error) {
	return s.Repo.ListBookingOfUsers(req)
}