package models

type Room struct {
    RoomNumber      int32   `json:"room_number"`
    RoomType        string  `json:"room_type"`
    PricePerNight   float32 `json:"price_per_night"`
    Availability    bool    `json:"availability"`
}

type Hotel struct {
    HotelID   string `json:"hotel_id"`
    Name      string `json:"name"`
    Location  string `json:"location"`
    Rating    int32  `json:"rating"`
    Address   string `json:"address"`
}

type ListOfHotelRequest struct{}

type ListOfHotelResponse struct {
    List []Hotel `json:"list"`
}

type GetDetailsOfHotelRequest struct {
    HotelID string `json:"hotel_id"`
}

type GetDetailsOfHotelResponse struct {
    HotelID   string `json:"hotel_id"`
    Name      string `json:"name"`
    Location  string `json:"location"`
    Rating    int32  `json:"rating"`
    Address   string `json:"address"`
    Rooms     []Room `json:"rooms"`
}

type GetAvailabilityRoomsRequest struct {
    HotelID string `json:"hotel_id"`
}

type AvailabilityRooms struct {
    RoomType       string `json:"room_type"`
    AvailableRooms int32  `json:"available_rooms"`
}

type GetAvailabilityRoomsResponse struct {
    List []AvailabilityRooms `json:"list"`
}
