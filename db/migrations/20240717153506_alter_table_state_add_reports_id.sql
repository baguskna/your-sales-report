-- +goose Up
-- +goose StatementBegin
ALTER TABLE stats ADD COLUMN reports_id UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stats DROP COLUMN reports_id;
-- +goose StatementEnd