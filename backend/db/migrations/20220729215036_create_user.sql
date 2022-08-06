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
-- password=password
INSERT INTO users (name, password, email) VALUES("test1", "$2a$10$s0IkhGD3R9qmwZ8/afJbP..uKGPNGl/ObUrVH8J2j181uk0KTfJ3q", "teat1@mail.com");
INSERT INTO users (name, password, email) VALUES("test2", "$2a$10$s0IkhGD3R9qmwZ8/afJbP..uKGPNGl/ObUrVH8J2j181uk0KTfJ3q", "teat2@mail.com");
INSERT INTO users (name, password, email) VALUES("test3", "$2a$10$s0IkhGD3R9qmwZ8/afJbP..uKGPNGl/ObUrVH8J2j181uk0KTfJ3q", "teat3@mail.com");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
