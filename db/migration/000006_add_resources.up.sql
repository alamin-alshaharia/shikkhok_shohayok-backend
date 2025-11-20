CREATE TABLE resources (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR NOT NULL,
  file_path VARCHAR NOT NULL,
  type VARCHAR NOT NULL, -- 'book', 'notice', etc.
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);
