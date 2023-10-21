-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO categories (name, slug) VALUES ('Docker', 'docker');
INSERT INTO categories (name, slug) VALUES ('Kubernetes', 'kubernetes');
INSERT INTO categories (name, slug) VALUES ('Golang', 'golang');
INSERT INTO categories (name, slug) VALUES ('Agile', 'agile');
INSERT INTO categories (name, slug) VALUES ('Requirement definition', 'requirement-definition');
INSERT INTO categories (name, slug) VALUES ('Nuxt', 'nuxt');

UPDATE articles SET category_id = 1 WHERE id = 1;
UPDATE articles SET category_id = 2 WHERE id = 2;
UPDATE articles SET category_id = 3 WHERE id = 3;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DELETE FROM categories WHERE 'name' = 'category-1';
DELETE FROM categories WHERE 'name' = 'category-2';
DELETE FROM categories WHERE 'name' = 'category-3';
-- +goose StatementEnd
