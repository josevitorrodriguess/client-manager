-- name: CreateUser :one
INSERT INTO users (name, email, password) 
VALUES ($1, $2, $3)
RETURNING id, name, email, password;


-- name: UpdateUser :one
UPDATE users
SET name = $1, email = $2, password = $3
WHERE id = $4
RETURNING id, name, email; 

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;



