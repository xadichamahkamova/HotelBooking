CREATE TABLE
    rooms (
        room_number SERIAL PRIMARY KEY,
        room_type VARCHAR(100) NOT NULL,
        price_per_night DECIMAL(10, 2) NOT NULL,
        availability BOOLEAN DEFAULT TRUE,
        hotel_id UUID REFERENCES hotels (id)
    );