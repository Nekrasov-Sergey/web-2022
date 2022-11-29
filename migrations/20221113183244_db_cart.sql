-- +goose Up
-- +goose StatementBegin
create table cart
(
    store    uuid REFERENCES stores (UUID) ON DELETE CASCADE,
    quantity int
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart;
-- +goose StatementEnd