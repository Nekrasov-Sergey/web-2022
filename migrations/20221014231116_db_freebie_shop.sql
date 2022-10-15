-- +goose Up
-- +goose StatementBegin
create table promos
(
    UUID     uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    Store    text,
    Discount text,
    Price    text,
    Quantity int,
    Promo    text[]
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE promos;
-- +goose StatementEnd