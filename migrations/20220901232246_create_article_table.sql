-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    thumbnail_url varchar(1024) NOT NULL,
    title varchar(255) NOT NULL,
    content text NOT NULL,
    status int,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE articles
    ADD FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE articles DROP CONSTRAINT articles_user_id_fkey;
DROP TABLE articles
-- +goose StatementEnd
