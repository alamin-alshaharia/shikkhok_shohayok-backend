-- name: CreateExam :one
INSERT INTO exams (
  class_id,
  name,
  total_marks
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ListExams :many
SELECT * FROM exams
WHERE class_id = $1
ORDER BY created_at DESC;

-- name: UpsertResult :one
INSERT INTO results (
  exam_id,
  student_id,
  obtained_marks
) VALUES (
  $1, $2, $3
)
ON CONFLICT (exam_id, student_id)
DO UPDATE SET obtained_marks = $3
RETURNING *;

-- name: GetResultsByExam :many
SELECT r.*, s.name as student_name, s.roll
FROM results r
JOIN students s ON r.student_id = s.id
WHERE r.exam_id = $1
ORDER BY s.roll;
