CREATE TABLE
    hotels (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name VARCHAR(255) NOT NULL,
        location VARCHAR(255) NOT NULL,
        rating INT CHECK (
            rating >= 0
            AND rating <= 5
        ),
        address TEXT
    );