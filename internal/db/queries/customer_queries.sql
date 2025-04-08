-- name: CreateCustomerPF :one
WITH new_customer AS (
    INSERT INTO customers (type, email, phone)
    VALUES ($1, $2, $3)
    RETURNING id
),
customer_pf AS (
    INSERT INTO customerf_pf (customer_id, cpf, name, birth_date)
    SELECT id, $4, $5, $6
    FROM new_customer
    RETURNING customer_id
)
INSERT INTO addresses (
    customer_id,
    address_type,
    street,
    number,
    complement,
    state,
    city,
    cep
)
SELECT 
    customer_id,
    $7, $8, $9, $10, $11, $12, $13
FROM customer_pf
RETURNING customer_id;


-- name: CreateCustomerPJ :one
WITH new_customer AS (
    INSERT INTO customers (type, email, phone)
    VALUES ($1, $2, $3)
    RETURNING id
),
customer_pj AS (
    INSERT INTO customerf_pj (customer_id, cnpj, company_name)
    SELECT id, $4, $5
    FROM new_customer
    RETURNING customer_id
)
INSERT INTO addresses (
    customer_id,
    address_type,
    street,
    number,
    complement,
    state,
    city,
    cep
)
SELECT 
    customer_id,
    $6, $7, $8, $9, $10, $11, $12
FROM customer_pj
RETURNING customer_id;

-- name: AddAddressToCustomer :one
INSERT INTO addresses (
    customer_id,
    address_type,
    street,
    number,
    complement,
    state,
    city,
    cep
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;


-- name: GetCustomerByID :one
SELECT 
    c.id,
    c.type,
    c.email,
    c.phone,
    c.is_active,
    c.created_at,
    c.updated_at,
    CASE 
        WHEN c.type = 'PF' THEN pf.cpf
        ELSE NULL
    END as cpf,
    CASE 
        WHEN c.type = 'PF' THEN pf.name
        ELSE NULL
    END as pf_name,
    CASE 
        WHEN c.type = 'PF' THEN pf.birth_date
        ELSE NULL
    END as birth_date,
    CASE 
        WHEN c.type = 'PJ' THEN pj.cnpj
        ELSE NULL
    END as cnpj,
    CASE 
        WHEN c.type = 'PJ' THEN pj.company_name
        ELSE NULL
    END as company_name
FROM customers c
LEFT JOIN customerf_pf pf ON c.id = pf.customer_id AND c.type = 'PF'
LEFT JOIN customerf_pj pj ON c.id = pj.customer_id AND c.type = 'PJ'
WHERE c.id = $1;



-- name: GetAllCustomers :many
SELECT
    c.id,
    c.email,
    c.phone,
    c.created_at,
    c.updated_at,
    c.is_active,
    pf.cpf,
    pf.name,
    pf.birth_date,
    pj.cnpj,
    pj.company_name,
    a.id AS address_id,
    a.address_type,
    a.street,
    a.number,
    a.complement,
    a.state,
    a.city,
    a.cep
FROM customers c
LEFT JOIN customerf_pf pf ON c.id = pf.customer_id
LEFT JOIN customerf_pj pj ON c.id = pj.customer_id
LEFT JOIN addresses a ON c.id = a.customer_id
ORDER BY c.id, a.id;


-- name: GetCustomerAddresses :many
SELECT 
    id,
    address_type,
    street,
    number,
    complement,
    state,
    city,
    cep
FROM addresses
WHERE customer_id = $1;


-- name: UpdateCustomerBasicInfo :one
UPDATE customers
SET email = $2, phone = $3
WHERE id = $1
RETURNING id;


-- name: UpdateAddress :one
UPDATE addresses
SET 
    address_type = $2,
    street = $3,
    number = $4,
    complement = $5,
    state = $6,
    city = $7,
    cep = $8
WHERE id = $1
RETURNING id;


-- name: DeleteAddress :exec
DELETE FROM addresses
WHERE id = $1;


















