-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXITS workouts (
    id BIGSERIAL PRIMARY KEY,
    -- user_id
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration_minutes INTEGER NOT NULL,
    calories_burned INTEGER,
    create_at TIMESTAMP WITH TIME ZONE DEFAULT CURENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workouts;
-- +goose StatementEnd
