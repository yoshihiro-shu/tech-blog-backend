-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid(),
    name varchar(256) NOT NULL,
    password varchar(256) NOT NULL,
    email varchar(256) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
