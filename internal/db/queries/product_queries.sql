-- name: GetProductByID :one
SELECT * FROM products
WHERE id = $1;

-- name: GetProductsByCustomerID :many
SELECT * FROM products
WHERE customer_id = $1
ORDER BY id;

-- name: GetUnpaidProductsByCustomerID :many
SELECT 
    id,
    customer_id,
    type_product,
    description,
    total_value,
    down_payment,
    is_paid,
    is_finished,
    (total_value - down_payment) AS remaining_to_pay
FROM products
WHERE customer_id = $1 AND is_paid = false
ORDER BY id;

-- name: GetUnfinishedProductsByCustomerID :many
SELECT * FROM products
WHERE customer_id = $1 AND is_finished = false
ORDER BY id;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: ListAllProducts :many
SELECT * FROM products
ORDER BY id;

-- name: UpdateProductPaymentStatus :one
UPDATE products
SET is_paid = $1
WHERE id = $2
RETURNING *;

-- name: UpdateProductFinishStatus :one
UPDATE products
SET is_finished = $1
WHERE id = $2
RETURNING *;

-- name: CountProductsByCustomerID :one
SELECT COUNT(*) FROM products
WHERE customer_id = $1;

-- name: GetProductsValueSumByCustomerID :one
SELECT SUM(total_value) FROM products
WHERE customer_id = $1;
