-- +goose Up
-- +goose StatementBegin
INSERT INTO articles (title, content, status) VALUES('test1', 'test-content-1', 2);
INSERT INTO articles (title, content, status) VALUES('test2', 'test-content-2', 2);
INSERT INTO articles (title, content, status) VALUES('test3', 'test-content-3', 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM articles WHERE title ='test1';
DELETE FROM articles WHERE title ='test2';
DELETE FROM articles WHERE title ='test3';
-- +goose StatementEnd
