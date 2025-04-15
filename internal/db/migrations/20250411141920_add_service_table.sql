-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    customer_id UUID NOT NULL REFERENCES customers(id),
    type_product VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    total_value NUMERIC(10, 2) NOT NULL,
    down_payment NUMERIC(10, 2) NOT NULL,
    is_paid BOOLEAN NOT NULL DEFAULT FALSE,
    is_finished BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS services;