-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    customer_id UUID REFERENCES customers(id),
    type_product VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    total_value DECIMAL(10, 2) NOT NULL,
    down_payment DECIMAL(10, 2) NOT NULL,
    is_paid BOOLEAN DEFAULT FALSE,
    is_finished BOOLEAN DEFAULT FALSE
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS services;