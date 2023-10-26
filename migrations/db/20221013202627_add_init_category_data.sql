-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO categories (name, slug) VALUES ('Agile', 'agile');
INSERT INTO categories (name, slug) VALUES ('Bussiness', 'bussiness');
INSERT INTO categories (name, slug) VALUES ('Marketing', 'marketing');
INSERT INTO categories (name, slug) VALUES ('Frontend', 'frontend');
INSERT INTO categories (name, slug) VALUES ('Backend', 'backend');
INSERT INTO categories (name, slug) VALUES ('Infrastructure', 'infrastructure');
INSERT INTO categories (name, slug) VALUES ('Requirement Definition', 'requirement-definition');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DELETE FROM categories WHERE 'name' = 'Agile';
DELETE FROM categories WHERE 'name' = 'Bussiness';
DELETE FROM categories WHERE 'name' = 'Backend';
DELETE FROM categories WHERE 'name' = 'Frontend';
DELETE FROM categories WHERE 'name' = 'Infrastructure';
DELETE FROM categories WHERE 'name' = 'Marketing';
DELETE FROM categories WHERE 'name' = 'System Design';
-- +goose StatementEnd
