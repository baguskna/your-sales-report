-- +goose Up
-- +goose StatementBegin
ALTER TABLE stats
ADD CONSTRAINT fk_role FOREIGN KEY (reports_id) REFERENCES reports (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stats DROP CONSTRAINT fk_role;
-- +goose StatementEnd