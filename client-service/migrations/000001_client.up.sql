CREATE TABLE clients (
    id UUID PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    age INT,
    gender VARCHAR(6),
    phone_number VARCHAR(15),
    address TEXT,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    status BOOLEAN NOT NULL DEFAULT TRUE,
    refresh TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT  CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ NULL
);
