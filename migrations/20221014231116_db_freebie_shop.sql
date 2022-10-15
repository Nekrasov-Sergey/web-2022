-- +goose Up
-- +goose StatementBegin
create table promos
(
    UUID     uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    Store    text NOT NULL,
    Discount text NOT NULL,
    Price    text NOT NULL,
    Quantity serial NOT NULL,
    Promo    text[] NOT NULL
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE promos;
-- +goose StatementEnd