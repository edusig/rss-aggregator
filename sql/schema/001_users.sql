-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR(200) NOT NULL
);

-- +goose Down
DROP TABLE users;