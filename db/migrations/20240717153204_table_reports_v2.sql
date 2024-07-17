-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reports (
    id UUID PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reports;
-- +goose StatementEnd