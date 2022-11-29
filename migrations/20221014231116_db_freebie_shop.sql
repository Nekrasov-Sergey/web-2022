-- +goose Up
-- +goose StatementBegin
create table stores
(
    uuid     uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    name     text,
    discount int,
    price    int,
    quantity int,
    promo    text[],
    image    text
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stores;
-- +goose StatementEnd