package repository

import (
	pb "booking-service/genproto/bookingpb"
)

type IBookingRepository interface {
	CreateBooking (req *pb.CreateBookingRequest) (*pb.Booking, error)
	GetDetailsOfBooking(req *pb.GetDetailsOfBookingRequest) (*pb.Booking, error)
	UpdateBooking(req *pb.UpdateBookingRequest) (*pb.Booking, error)
	CancelBooking(req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error)
	ListBookingOfUsers(req *pb.ListBookingOfUserRequest) (*pb.ListBookingOfUserResponse, error)
}