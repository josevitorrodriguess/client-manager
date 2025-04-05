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


-- name: GetCustomerByID :one
SELECT * FROM customers
WHERE id = $1;


-- name: GetCustomerByEmail :one
SELECT * FROM customers
WHERE email = $1;


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


-- -- name: UpdateCustomer :one
-- UPDATE customers
-- SET email = $2, phone = $3
-- WHERE id = $1
-- RETURNING id, email, phone;


-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;


-- name: GetAllCustomers :many
SELECT * FROM customers;


-- name: GetCustomerAddresses :many
SELECT * FROM addresses
WHERE customer_id = $1;


-- name: UpdateCustomerPF :one
UPDATE customerf_pf
SET name = $2, cpf = $3, birth_date = $4
WHERE customer_id = $1
RETURNING customer_id;


-- name: UpdateCustomerPJ :one
UPDATE customerf_pj
SET company_name = $2, cnpj = $3
WHERE customer_id = $1
RETURNING customer_id;


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


-- name: GetCustomerPFDetails :one
SELECT 
    c.id,
    c.email,
    c.phone,
    c.created_at,
    c.updated_at,
    c.is_active,
    pf.cpf,
    pf.name,
    pf.birth_date
FROM customers c
JOIN customerf_pf pf ON c.id = pf.customer_id
WHERE c.id = $1;


-- name: GetCustomerPJDetails :one
SELECT 
    c.id,
    c.email,
    c.phone,
    c.created_at,
    c.updated_at,
    c.is_active,
    pj.cnpj,
    pj.company_name
FROM customers c
JOIN customerf_pj pj ON c.id = pj.customer_id
WHERE c.id = $1;


-- name: ListActiveCustomers :many
SELECT * FROM customers
WHERE is_active = true
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;


-- name: SetCustomerStatus :exec
UPDATE customers
SET is_active = $2
WHERE id = $1;


-- name: SearchCustomersByEmail :many
SELECT * FROM customers
WHERE email ILIKE $1
LIMIT $2;



-- name: SearchPJCustomersByCompanyName :many
SELECT c.* FROM customers c
JOIN customerf_pj pj ON c.id = pj.customer_id
WHERE pj.company_name ILIKE $1
LIMIT $2;


-- name: DeleteAddress :exec
DELETE FROM addresses
WHERE id = $1;


-- name: SearchPFCustomersByName :many
SELECT c.* FROM customers c
JOIN customerf_pf pf ON c.id = pf.customer_id
WHERE pf.name ILIKE $1
LIMIT $2;


-- name: GetRecentCustomers :many
SELECT * FROM customers
ORDER BY created_at DESC
LIMIT $1;













