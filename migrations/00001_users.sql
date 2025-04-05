-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXITS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    bio TEXT,
    create_at TIMESTAMP WITH TIME ZONE DEFAULT CURENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
