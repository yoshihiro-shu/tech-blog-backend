-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles (
    id SERIAL,
    title varchar(256) NOT NULL,
    content text NOT NULL,
    status int,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE articles
-- +goose StatementEnd
