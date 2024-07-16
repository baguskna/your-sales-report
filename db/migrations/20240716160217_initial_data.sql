-- +goose Up
-- +goose StatementBegin
INSERT INTO
    reports (id, total, date)
VALUES (
        uuid_generate_v4 (),
        14126,
        '2024-07-14'
    ),
    (
        uuid_generate_v4 (),
        14137,
        '2024-07-16'
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM reports WHERE date IN ('2024-07-14', '2024-07-16');
-- +goose StatementEnd