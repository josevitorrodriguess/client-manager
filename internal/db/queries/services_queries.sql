-- name: CreateService :one
INSERT INTO services (
    customer_id,
    type_product,
    description,
    total_value,
    down_payment,
    is_paid,
    is_finished
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) 
RETURNING id;

-- name: GetServicesByCustomerID :many
SELECT * FROM services
WHERE customer_id = $1
ORDER BY id;

-- name: DeleteService :exec
DELETE FROM services
WHERE id = $1;

-- name: ListAllServices :many
SELECT * FROM services
ORDER BY id;

-- name: UpdateServicePaymentStatus :one
UPDATE services
SET is_paid = $1
WHERE id = $2
RETURNING *;

-- name: UpdateServiceFinishStatus :one
UPDATE services
SET is_finished = $1
WHERE id = $2
RETURNING *; 

-- name: CountServicesByCustomerID :one
SELECT COUNT(*) FROM services
WHERE customer_id = $1;
