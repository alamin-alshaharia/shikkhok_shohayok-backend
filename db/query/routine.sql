-- name: CreateRoutine :one
INSERT INTO routines (
  class_id,
  day_of_week,
  start_time,
  end_time,
  subject,
  teacher_name
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: ListRoutinesByClass :many
SELECT * FROM routines
WHERE class_id = $1
ORDER BY
  CASE
    WHEN day_of_week = 'Sunday' THEN 1
    WHEN day_of_week = 'Monday' THEN 2
    WHEN day_of_week = 'Tuesday' THEN 3
    WHEN day_of_week = 'Wednesday' THEN 4
    WHEN day_of_week = 'Thursday' THEN 5
    WHEN day_of_week = 'Friday' THEN 6
    WHEN day_of_week = 'Saturday' THEN 7
  END,
  start_time;

-- name: CreateLessonPlan :one
INSERT INTO lesson_plans (
  class_id,
  date,
  topic,
  note
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: ListLessonPlans :many
SELECT * FROM lesson_plans
WHERE class_id = $1
ORDER BY date DESC;

-- name: GetTeacherDailyRoutinesCount :one
SELECT COUNT(*)
FROM routines r
JOIN classes c ON r.class_id = c.id
WHERE c.teacher_id = $1 AND r.day_of_week = $2;
