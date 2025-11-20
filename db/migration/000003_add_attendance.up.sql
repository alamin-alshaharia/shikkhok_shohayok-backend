CREATE TABLE attendance (
  id BIGSERIAL PRIMARY KEY,
  student_id BIGINT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
  date DATE NOT NULL DEFAULT CURRENT_DATE,
  status VARCHAR NOT NULL, -- 'present', 'absent', 'late'
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  UNIQUE(student_id, date)
);
