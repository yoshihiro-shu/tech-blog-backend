-- +goose Up
-- +goose StatementBegin
INSERT INTO articles (user_id, title, content, status) VALUES(1, 'test1', 'test-content-1', 2);
INSERT INTO articles (user_id, title, content, status) VALUES(1, 'test2', 'test-content-2', 2);
INSERT INTO articles (user_id, title, content, status) VALUES(1, 'test3', 'test-content-3', 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM articles WHERE title = 'test1';
DELETE FROM articles WHERE title = 'test2';
DELETE FROM articles WHERE title = 'test3';
-- +goose StatementEnd
