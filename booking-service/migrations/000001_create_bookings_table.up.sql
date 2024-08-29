CREATE TYPE booking_status AS ENUM ('PENDING', 'CONFIRMED', 'CANCELLED');

CREATE TABLE
    bookings (
        booking_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID REFERENCES users (id),
        hotel_id UUID REFERENCES hotels (id),
        room_type VARCHAR(100) NOT NULL,
        check_in_date DATE NOT NULL,
        check_out_date DATE NOT NULL,
        total_amount NUMERIC(10, 2) NOT NULL,
        status booking_status NOT NULL DEFAULT 'PENDING'
    );