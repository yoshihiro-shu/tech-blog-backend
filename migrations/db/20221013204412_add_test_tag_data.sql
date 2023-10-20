-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO tags (name, slug) VALUES ('Docker', 'docker');
INSERT INTO tags (name, slug) VALUES ('Kubernetes', 'kubernetes');
INSERT INTO tags (name, slug) VALUES ('Golang', 'golang');
INSERT INTO tags (name, slug) VALUES ('Agile', 'agile');
INSERT INTO tags (name, slug) VALUES ('Requirement definition', 'requirement-definition');
INSERT INTO tags (name, slug) VALUES ('Nuxt', 'nuxt');

INSERT INTO article_tags (article_id, tag_id) VALUES(1, 1);
INSERT INTO article_tags (article_id, tag_id) VALUES(1, 2);
INSERT INTO article_tags (article_id, tag_id) VALUES(2, 2);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DELETE FROM articles WHERE 'name' = 'tag-1';
DELETE FROM articles WHERE 'name' = 'tag-2';
DELETE FROM articles WHERE 'name' = 'tag-3';
-- +goose StatementEnd
