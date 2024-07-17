-- +goose Up
-- +goose StatementBegin
INSERT INTO
    reports (id, title)
VALUES (
        uuid_generate_v4 (),
        'Tanamera Coffee Drip Bag / Filter Bag: Breakfast Blend'
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM reports
WHERE
    title = 'Tanamera Coffee Drip Bag / Filter Bag: Breakfast Blend';
-- +goose StatementEnd