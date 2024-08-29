package repository

import (
	"database/sql"
	pb "hotel-service/genproto/hotelpb"
	"hotel-service/logger"
)

type HotelRepo struct {
	DB *sql.DB
}

func NewHotelRepository(db *sql.DB) IHotelRepository {
	return &HotelRepo{
		DB: db,
	}
}

func (db *HotelRepo) ListOfHotel() (*pb.ListOfHotelResponse, error) {

	logger.Info("ListOfHotel started..")
	resp := pb.ListOfHotelResponse{}
	query := `
	SELECT * 
	FROM hotels`
	rows, err := db.DB.Query(query)
	if err != nil {
		logger.Error("Failed to query hotels:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := pb.Hotel{}
		err := rows.Scan(
			&item.HotelId,
			&item.Name,
			&item.Location,
			&item.Rating,
			&item.Address,
		)
		if err != nil {
			logger.Error("Failed to scan hotel row:", err)
			return nil, err
		}
		resp.List = append(resp.List, &item)
	}
	logger.Info("ListOfHotel successfully fetched hotels")
	return &resp, nil
}

func (db *HotelRepo) GetDetailsOfHotel(req *pb.GetDetailsOfHotelRequest) (*pb.GetDetailsOfHotelResponse, error) {

	logger.Info("GetDetailsOfHotel called with hotel ID:", req.HotelId)
	resp := pb.GetDetailsOfHotelResponse{}
	rooms := []*pb.Room{}
	query := `
	SELECT h.id, h.name, h.location, h.rating, h.address, r.room_number, r.room_type, r.price_per_night, r.availability
	FROM hotels as h
	INNER JOIN rooms as r 
	ON h.id = r.hotel_id
	WHERE id=$1`
	rows, err := db.DB.Query(query, req.HotelId)
	if err != nil {
		logger.Error("Failed to query hotel details:", err)	
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		room := pb.Room{}
		err := rows.Scan(
			&resp.HotelId,
			&resp.Name,
			&resp.Location,
			&resp.Rating,
			&resp.Address,
			&room.RoomNumber,
			&room.RoomType,
			&room.PricePerNight,
			&room.Availability,
		)
		if err != nil {
			logger.Error("Failed to scan hotel and room details:", err)
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	resp.Rooms = rooms
	logger.Info("GetDetailsOfHotel successfully fetched details for hotel ID:", req.HotelId)
	return &resp, nil
}

func (db *HotelRepo) GetAvailabilityRooms(req *pb.GetAvailabilityRoomsRequest) (*pb.GetAvailabilityRoomsResponse, error) {

	logger.Info("GetAvailabilityRooms called with hotel ID:", req.HotelId)
	resp := pb.GetAvailabilityRoomsResponse{}
	query := `
	SELECT room_type, room_number 
	FROM rooms
	WHERE hotel_id=$1 AND availability=true
	`
	rows, err := db.DB.Query(query, req.HotelId)
	if err != nil {
		logger.Error("Failed to query available rooms:", err)
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		room := pb.AvailabilityRooms{}
		err := rows.Scan(
			&room.RoomType,
			&room.AvailableRooms,
		)
		if err != nil {
			logger.Error("Failed to scan available room:", err)
			return nil, err
		}
		resp.List = append(resp.List, &room)
	}
	logger.Info("GetAvailabilityRooms successfully fetched available rooms for hotel ID:", req.HotelId)
	return &resp, nil
}
