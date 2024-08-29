package repository

import (
	pb "booking-service/genproto/bookingpb"
	"booking-service/logger"
	"database/sql"
	"log"
)

type BookingRepo struct {
	DB *sql.DB
}

func NewBookingRepository(db *sql.DB) IBookingRepository {
	return &BookingRepo{
		DB: db,
	}
}

func (db *BookingRepo) CreateBooking(req *pb.CreateBookingRequest) (*pb.Booking, error) {

	logger.Info("CreateBooking called with request:", req)
	log.Println(req)
	resp := pb.Booking{}
	query := `
	INSERT INTO bookings(user_id, hotel_id, room_type, check_in_date, check_out_date, total_amount)
	VALUES($1, $2, $3, $4, $5, $6) 
	RETURNING booking_id, user_id, hotel_id, room_type, check_in_date, check_out_date, total_amount, status`
	err := db.DB.QueryRow(query,
		req.UserId,
		req.HotelId,
		req.RoomType,
		req.CheckInDate,
		req.CheckOutDate,
		req.TotalAmount,
	).Scan(
		&resp.BookingId,
		&resp.UserId,
		&resp.HotelId,
		&resp.RoomType,
		&resp.CheckInDate,
		&resp.CheckOutDate,
		&resp.TotalAmount,
		&resp.Status,
	)
	if err != nil {
		logger.Error("Error creating booking:", err)
		return nil, err
	}

	logger.Info("Booking created successfully with ID:", resp.BookingId)
	return &resp, nil
}

func (db *BookingRepo) GetDetailsOfBooking(req *pb.GetDetailsOfBookingRequest) (*pb.Booking, error) {

	logger.Info("GetDetailsOfBooking called with request:", req)

	resp := pb.Booking{}
	query := `
	SELECT * 
	FROM bookings 
	WHERE booking_id = $1`
	err := db.DB.QueryRow(query,
		req.BookingId,
	).Scan(
		&resp.BookingId,
		&resp.UserId,
		&resp.HotelId,
		&resp.RoomType,
		&resp.CheckInDate,
		&resp.CheckOutDate,
		&resp.TotalAmount,
		&resp.Status,
	)
	if err != nil {
		logger.Error("Error getting booking details:", err)
		return nil, err
	}
	logger.Info("Booking details retrieved successfully for ID:", resp.BookingId)
	return &resp, nil
}

func (db *BookingRepo) UpdateBooking(req *pb.UpdateBookingRequest) (*pb.Booking, error) {

	logger.Info("UpdateBooking called with request:", req)

	resp := pb.Booking{}
	query := `
	UPDATE bookings 
	SET check_in_date=$1, check_out_date=$2, total_amount=$3, status=$4
	WHERE booking_id=$5
	RETURNING booking_id, user_id, hotel_id, room_type, check_in_date, check_out_date, total_amount, status
	`
	err := db.DB.QueryRow(query,
		req.CheckInDate,
		req.CheckOutDate,
		req.TotalAmount,
		req.Status,
		req.BookingId,
	).Scan(
		&resp.BookingId,
		&resp.UserId,
		&resp.HotelId,
		&resp.RoomType,
		&resp.CheckInDate,
		&resp.CheckOutDate,
		&resp.TotalAmount,
		&resp.Status,
	)
	if err != nil {
		logger.Error("Error updating booking:", err)
		return nil, err
	}
	logger.Info("Booking updated successfully with ID:", resp.BookingId)
	return &resp, nil
}

func (db *BookingRepo) CancelBooking(req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error) {

	logger.Info("CancelBooking called with request:", req)

	resp := pb.CancelBookingResponse{}
	query := `
	UPDATE bookings 
	SET status = 'CANCELLED'
	WHERE booking_id = $1
	RETURNING booking_id;
	`
	err := db.DB.QueryRow(query, req.BookingId).Scan(&resp.BookingId)
	if err != nil {
		logger.Error("Error cancelling booking:", err)
		return nil, err
	}
	resp.Message = "Booking cancelled succesfully"
	logger.Info("Booking cancelled successfully with ID:", resp.BookingId)
	return &resp, nil
}

func (db *BookingRepo) ListBookingOfUsers(req *pb.ListBookingOfUserRequest) (*pb.ListBookingOfUserResponse, error) {

	logger.Info("ListBookingOfUsers called with request:", req)

	resp := pb.ListBookingOfUserResponse{}
	query := `
	SELECT booking_id, hotel_id, room_type, check_in_date, check_out_date, total_amount, status
	FROM bookings 
	WHERE user_id=$1`
	rows, err := db.DB.Query(query, req.UserId)
	if err != nil {
		logger.Error("Error listing bookings for user:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		booking := pb.UsersBooking{}
		err := rows.Scan(
			&booking.BookingId,
			&booking.HotelId,
			&booking.RoomType,
			&booking.CheckInDate,
			&booking.CheckOutDate,
			&booking.TotalAmount,
			&booking.Status,
		)
		if err != nil {
			logger.Error("Error scanning booking:", err)
			return nil, err
		}
		resp.List = append(resp.List, &booking)
	}
	logger.Info("Bookings listed successfully for user ID:", req.UserId)
	return &resp, nil
}
