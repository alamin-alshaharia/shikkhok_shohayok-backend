-- name: CreateUser :one
INSERT INTO users (
  phone,
  password,
  full_name,
  school_name
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = $1 LIMIT 1;
