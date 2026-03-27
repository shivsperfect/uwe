-- +goose Up
CREATE TABLE customers (
    id UUID NOT NULL PRIMARY KEY,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

-- +goose Down
DROP TABLE customers;
