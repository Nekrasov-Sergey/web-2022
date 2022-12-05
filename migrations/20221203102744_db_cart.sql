-- +goose Up
-- +goose StatementBegin
create table cart
(
    store_uuid uuid REFERENCES stores (uuid) ON DELETE CASCADE,
    user_uuid  uuid REFERENCES users (uuid) ON DELETE CASCADE,
    quantity   int
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart;
-- +goose StatementEnd