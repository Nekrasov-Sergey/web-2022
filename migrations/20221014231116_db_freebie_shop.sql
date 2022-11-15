-- +goose Up
-- +goose StatementBegin
create table stores
(
    UUID     uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    Name     text,
    Discount int,
    Price    int,
    Quantity int,
    Promo    text[],
    Image    text
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stores;
-- +goose StatementEnd