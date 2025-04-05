-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    customer_id UUID NOT NULL REFERENCES customers(id),
    address_type VARCHAR(50) NOT NULL,
    street VARCHAR(255) NOT NULL,
    number VARCHAR(50) NOT NULL,
    complement VARCHAR(255),
    state VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    cep VARCHAR(10) NOT NULL
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS addresses;
