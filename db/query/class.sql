-- name: CreateClass :one
INSERT INTO classes (
  name,
  section,
  teacher_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetClass :one
SELECT * FROM classes
WHERE id = $1 LIMIT 1;

-- name: ListClasses :many
SELECT * FROM classes
WHERE teacher_id = $1
ORDER BY id;

-- name: UpdateClass :one
UPDATE classes
SET name = $2, section = $3
WHERE id = $1
RETURNING *;

-- name: DeleteClass :exec
DELETE FROM classes
WHERE id = $1;
