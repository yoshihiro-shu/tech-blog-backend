-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO categories (name, slug) VALUES ('Agile', 'agile');
INSERT INTO categories (name, slug) VALUES ('Bussiness', 'bussiness');
INSERT INTO categories (name, slug) VALUES ('Backend', 'backend');
INSERT INTO categories (name, slug) VALUES ('Frontend', 'frontend');
INSERT INTO categories (name, slug) VALUES ('Infrastructure', 'infrastructure');
INSERT INTO categories (name, slug) VALUES ('Marketing', 'marketing');

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
