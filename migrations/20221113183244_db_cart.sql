-- +goose Up
-- +goose StatementBegin
create table cart
(
    Store    uuid REFERENCES stores (UUID) ON DELETE CASCADE,
    Quantity int
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart;
-- +goose StatementEnd