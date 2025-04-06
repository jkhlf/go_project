-- +goose Up
-- +goose StatementBegin
-- First add the column without the NOT NULL constraint
ALTER TABLE workouts 
ADD COLUMN user_id BIGINT REFERENCES users(id) ON DELETE CASCADE;

-- Update existing workouts to have a default user_id (using the first user in the system)
-- You may want to adjust this to use a specific user_id that makes sense for your data
UPDATE workouts
SET user_id = (SELECT id FROM users ORDER BY id LIMIT 1);

-- Now add the NOT NULL constraint
ALTER TABLE workouts
ALTER COLUMN user_id SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE workouts DROP COLUMN user_id;
-- +goose StatementEnd