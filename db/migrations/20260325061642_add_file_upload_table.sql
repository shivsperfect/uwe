-- +goose Up
CREATE TABLE file_uploads (
    id uuid NOT NULL PRIMARY KEY,
    customer_id uuid REFERENCES customers,
    type int NOT NULL ,
    mapping jsonb default '{}',
    created_at timestamp default now(),
    updated_at timestamp,
    deleted_at timestamp
);

-- +goose Down
DROP TABLE file_uploads;
