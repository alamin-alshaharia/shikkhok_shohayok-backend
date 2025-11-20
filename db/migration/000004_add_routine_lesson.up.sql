CREATE TABLE routines (
  id BIGSERIAL PRIMARY KEY,
  class_id BIGINT NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  day_of_week VARCHAR NOT NULL, -- Sunday, Monday, etc.
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  subject VARCHAR NOT NULL,
  teacher_name VARCHAR NOT NULL, -- Optional override or additional info
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE lesson_plans (
  id BIGSERIAL PRIMARY KEY,
  class_id BIGINT NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  date DATE NOT NULL,
  topic VARCHAR NOT NULL,
  note TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);
