-- +goose Up
-- +goose StatementBegin
CREATE TABLE reports (
    id UUID PRIMARY KEY NOT NULL,
    total INT NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reports;
-- +goose StatementEnd