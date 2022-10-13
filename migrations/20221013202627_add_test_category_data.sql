-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO categories (name, description) VALUES('category-1', 'it is test category-1');
INSERT INTO categories (name, description) VALUES('category-2', 'it is test category-2');
INSERT INTO categories (name, description) VALUES('category-3', 'it is test category-3');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DELETE FROM categories WHERE name = 'category-1';
DELETE FROM categories WHERE name = 'category-2';
DELETE FROM categories WHERE name = 'category-3';
-- +goose StatementEnd
