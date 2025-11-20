-- name: CreateAttendance :one
INSERT INTO attendance (
  student_id,
  date,
  status
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAttendanceByDate :many
SELECT a.*, s.name as student_name, s.roll
FROM attendance a
JOIN students s ON a.student_id = s.id
WHERE s.class_id = $1 AND a.date = $2;

-- name: GetAttendanceReport :many
SELECT a.*, s.name as student_name, s.roll
FROM attendance a
JOIN students s ON a.student_id = s.id
WHERE s.class_id = $1 AND a.date BETWEEN $2 AND $3;

-- name: UpsertAttendance :one
INSERT INTO attendance (student_id, date, status)
VALUES ($1, $2, $3)
ON CONFLICT (student_id, date)
DO UPDATE SET status = $3
RETURNING *;

-- name: GetTeacherDailyAttendanceStats :one
SELECT
    COUNT(*) FILTER (WHERE a.status = 'present') as present_count,
    COUNT(*) as total_count
FROM attendance a
JOIN students s ON a.student_id = s.id
JOIN classes c ON s.class_id = c.id
WHERE c.teacher_id = $1 AND a.date = $2;
