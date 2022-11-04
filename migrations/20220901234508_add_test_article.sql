-- +goose Up
-- +goose StatementBegin
INSERT INTO articles (user_id, thumbnail_url, title, content, status) VALUES(1, 'https://source.unsplash.com/collection/1346951/1000x500?sig=1', 'test1', 'test-content-1', 2);
INSERT INTO articles (user_id, thumbnail_url, title, content, status) VALUES(1, 'https://source.unsplash.com/collection/1346951/1000x500?sig=2', 'test2', 'test-content-2', 2);
INSERT INTO articles (user_id, thumbnail_url, title, content, status) VALUES(1, 'https://source.unsplash.com/collection/1346951/1000x500?sig=3', 'test3', 'test-content-3', 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM articles WHERE title = 'test1';
DELETE FROM articles WHERE title = 'test2';
DELETE FROM articles WHERE title = 'test3';
-- +goose StatementEnd
