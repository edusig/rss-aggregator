-- +goose Up
CREATE TABLE posts (
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    title VARCHAR(400) NOT NULL,
    url VARCHAR(400) NOT NULL UNIQUE,
    description TEXT,
    published_at TIMESTAMP WITH TIME ZONE,
    feed_id UUID NOT NULL,
    FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts