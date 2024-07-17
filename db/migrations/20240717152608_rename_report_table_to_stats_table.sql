-- +goose Up
-- +goose StatementBegin
ALTER TABLE reports RENAME TO stats;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stats RENAME to reports;
-- +goose StatementEnd