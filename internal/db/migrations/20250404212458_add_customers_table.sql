-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- Criar o tipo enum personalizado
CREATE TYPE customer_type AS ENUM ('PF', 'PJ');

CREATE TABLE customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type customer_type NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT NOT NULL UNIQUE,
        
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);


CREATE TABLE customerf_pf(
    customer_id UUID PRIMARY KEY REFERENCES customers(id),
    cpf VARCHAR(14) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TABLE customerf_pj(
    customer_id UUID PRIMARY KEY REFERENCES customers(id),
    cnpj VARCHAR(14) NOT NULL UNIQUE,
    company_name VARCHAR(255) NOT NULL,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS customerf_pj;
DROP TABLE IF EXISTS customerf_pf;
DROP TABLE IF EXISTS customers;
DROP TYPE IF EXISTS customer_type;
