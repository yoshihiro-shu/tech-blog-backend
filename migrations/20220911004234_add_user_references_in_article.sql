-- +goose Up
-- +goose StatementBegin
ALTER TABLE articles
    ADD user_id int
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE articles DROP COLUMN user_id CASCADE;
-- +goose StatementEnd
