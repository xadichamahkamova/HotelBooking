package models

type Booking struct {
    BookingID    string  `json:"booking_id"`
    UserID       string  `json:"user_id"`
    HotelID      string  `json:"hotel_id"`
    RoomType     string  `json:"room_type"`
    CheckInDate  string  `json:"check_in_date"`
    CheckOutDate string  `json:"check_out_date"`
    TotalAmount  float32 `json:"total_amount"`
    Status       string  `json:"status"`
}

type CreateBookingRequest struct {
    UserID       string  `json:"user_id"`
    HotelID      string  `json:"hotel_id"`
    RoomType     string  `json:"room_type"`
    CheckInDate  string  `json:"check_in_date"`
    CheckOutDate string  `json:"check_out_date"`
    TotalAmount  float32 `json:"total_amount"`
}

type GetDetailsOfBookingRequest struct {
    BookingID string `json:"booking_id"`
}

type UpdateBookingRequest struct {
    BookingID    string  `json:"booking_id"`
    CheckInDate  string  `json:"check_in_date"`
    CheckOutDate string  `json:"check_out_date"`
    TotalAmount  float32 `json:"total_amount"`
    Status       string  `json:"status"`
}

type CancelBookingRequest struct {
    BookingID string `json:"booking_id"`
}

type CancelBookingResponse struct {
    Message   string `json:"message"`
    BookingID string `json:"booking_id"`
}

type ListBookingOfUserRequest struct {
    UserID string `json:"user_id"`
}

type UsersBooking struct {
    UserID       string  `json:"user_id"`
    BookingID    string  `json:"booking_id"`
    HotelID      string  `json:"hotel_id"`
    RoomType     string  `json:"room_type"`
    CheckInDate  string  `json:"check_in_date"`
    CheckOutDate string  `json:"check_out_date"`
    TotalAmount  float32 `json:"total_amount"`
    Status       string  `json:"status"`
}

type ListBookingOfUserResponse struct {
    List []UsersBooking `json:"list"`
}
