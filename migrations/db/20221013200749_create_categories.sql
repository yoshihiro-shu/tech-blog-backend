-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name varchar(255) NOT NULL,
  slug varchar(255) NOT NULL,
  description varchar(255),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE articles
  ADD COLUMN category_id INTEGER;

ALTER TABLE articles
  ADD FOREIGN KEY (category_id)
  REFERENCES categories (id)
  ON DELETE SET NULL;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE articles DROP CONSTRAINT articles_category_id_fkey;
ALTER TABLE articles DROP COLUMN category_id;

DROP TABLE categories;
-- +goose StatementEnd
