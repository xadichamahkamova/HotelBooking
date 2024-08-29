package service

import (
	"context"

	pbBooking "api-gateway/genproto/bookingpb"
	pbHotel "api-gateway/genproto/hotelpb"
	pbNotif "api-gateway/genproto/notificationpb"
	pbUser "api-gateway/genproto/userpb"
)

type ServiceRepositoryClient struct {
	userClient         pbUser.UserServiceClient
	hotelClient        pbHotel.HotelServiceClient
	bookingClient      pbBooking.BookingServiceClient
	notificationClienr pbNotif.NotificationServiceClient
}

func NewServiceRepositoryClient(
	conn1 *pbUser.UserServiceClient,
	conn2 *pbHotel.HotelServiceClient,
	conn3 *pbBooking.BookingServiceClient,
	conn4 *pbNotif.NotificationServiceClient,
) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		userClient:         *conn1,
		hotelClient:        *conn2,
		bookingClient:      *conn3,
		notificationClienr: *conn4,
	}
}

// User methods
func (s *ServiceRepositoryClient) RegisterUser(ctx context.Context, req *pbUser.RegisterUserRequest) (*pbUser.RegisterUserResponse, error) {
	return s.userClient.RegisterUser(ctx, req)
}

func (s *ServiceRepositoryClient) LoginUser(ctx context.Context, req *pbUser.LoginUserRequest) (*pbUser.LoginUserResponse, error) {
	return s.userClient.LoginUser(ctx, req)
}

// Hotel methods
func (s *ServiceRepositoryClient) ListOfHotel(ctx context.Context, req *pbHotel.ListOfHotelRequest) (*pbHotel.ListOfHotelResponse, error) {
	return s.hotelClient.ListOfHotel(ctx, req)
}

func (s *ServiceRepositoryClient) GetDetailsOfHotel(ctx context.Context, req *pbHotel.GetDetailsOfHotelRequest) (*pbHotel.GetDetailsOfHotelResponse, error) {
	return s.hotelClient.GetDetailsOfHotel(ctx, req)
}

func (s *ServiceRepositoryClient) GetAvailabilityRooms(ctx context.Context, req *pbHotel.GetAvailabilityRoomsRequest) (*pbHotel.GetAvailabilityRoomsResponse, error) {
	return s.hotelClient.GetAvailabilityRooms(ctx, req)
}

// Booking methods
func (s *ServiceRepositoryClient) CreateBooking(ctx context.Context, req *pbBooking.CreateBookingRequest) (*pbBooking.Booking, error) {
	return s.bookingClient.CreateBooking(ctx, req)
}

func (s *ServiceRepositoryClient) GetDetailsOfBooking(ctx context.Context, req *pbBooking.GetDetailsOfBookingRequest) (*pbBooking.Booking, error) {
	return s.bookingClient.GetDetailsOfBooking(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateBooking(ctx context.Context, req *pbBooking.UpdateBookingRequest) (*pbBooking.Booking, error) {
	return s.bookingClient.UpdateBooking(ctx, req)
}

func (s *ServiceRepositoryClient) CancelBooking(ctx context.Context, req *pbBooking.CancelBookingRequest) (*pbBooking.CancelBookingResponse, error) {
	return s.bookingClient.CancelBooking(ctx, req)
}

func (s *ServiceRepositoryClient) ListBookingOfUsers(ctx context.Context, req *pbBooking.ListBookingOfUserRequest) (*pbBooking.ListBookingOfUserResponse, error) {
	return s.bookingClient.ListBookingOfUsers(ctx, req)
}


// Notification methods 
func (s *ServiceRepositoryClient) SendEmail(ctx context.Context, req *pbNotif.SendEmailRequest) (*pbNotif.SendEmailResponse, error) {
	return s.notificationClienr.SendEmail(ctx, req)
} 