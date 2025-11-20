CREATE TABLE exams (
  id BIGSERIAL PRIMARY KEY,
  class_id BIGINT NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  name VARCHAR NOT NULL, -- e.g., "Midterm 2024"
  total_marks INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE results (
  id BIGSERIAL PRIMARY KEY,
  exam_id BIGINT NOT NULL REFERENCES exams(id) ON DELETE CASCADE,
  student_id BIGINT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
  obtained_marks INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  UNIQUE(exam_id, student_id)
);
