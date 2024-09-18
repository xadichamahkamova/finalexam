CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_name VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(150) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE
);
