-- name: CreateResource :one
INSERT INTO resources (
  title,
  file_path,
  type
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ListResources :many
SELECT * FROM resources
ORDER BY created_at DESC;
