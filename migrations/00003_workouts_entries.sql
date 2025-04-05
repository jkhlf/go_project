-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXITS workout_entries (
    id BIGSERIAL PRIMARY KEY,
    workout_id BIGINT NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
    exercise_name VARCHAR(255) NOT NULL,
    sets INTEGE NOT NULL,
    reps INTEGER,
    duration_second INTEGER,
    weight DECIMAL(5, 2),
    notes TEXT,
    order_index INTEGER NOT NULL,
    create_at TIMESTAMP WITH TIME ZONE DEFAULT CURENT_TIMESTAMP,
    CONSTRAINT valid_workout_entry CHECK(
    (reps IS NOT NULL OR duration_second IS NOT NULL) AND 
    (reps IS NULL OR duration_second IS NULL)
    )
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workout_entries;
-- +goose StatementEnd
