-- name: CreateUser :one
INSERT INTO users (name, email, password) 
VALUES ($1, $2, $3)
RETURNING id;


-- name: UpdateUser :one
UPDATE users
SET name = $1, email = $2, password = $3
WHERE id = $4
RETURNING id, name, email; 


-- name: GetUserByEmail :one
SELECT 
    id,
    name,
    email,
    password,
    created_at,
    updated_at
FROM users
WHERE email = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;



