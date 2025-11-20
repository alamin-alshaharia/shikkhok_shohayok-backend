-- name: CreateStudent :one
INSERT INTO students (
  name,
  roll,
  father_name,
  mobile_no,
  class_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetStudent :one
SELECT * FROM students
WHERE id = $1 LIMIT 1;

-- name: ListStudents :many
SELECT * FROM students
WHERE class_id = $1
ORDER BY roll;

-- name: UpdateStudent :one
UPDATE students
SET name = $2, roll = $3, father_name = $4, mobile_no = $5
WHERE id = $1
RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM students
WHERE id = $1;
