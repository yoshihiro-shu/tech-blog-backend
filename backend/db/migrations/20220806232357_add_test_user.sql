-- +goose Up
-- +goose StatementBegin
-- password=password
INSERT INTO users (name, password, email) VALUES('test1', '$2a$10$s0IkhGD3R9qmwZ8/afJbP..uKGPNGl/ObUrVH8J2j181uk0KTfJ3q', 'teat1@mail.com');
INSERT INTO users (name, password, email) VALUES('test2', '$2a$10$s0IkhGD3R9qmwZ8/afJbP..uKGPNGl/ObUrVH8J2j181uk0KTfJ3q', 'teat2@mail.com');
INSERT INTO users (name, password, email) VALUES('test3', '$2a$10$s0IkhGD3R9qmwZ8/afJbP..uKGPNGl/ObUrVH8J2j181uk0KTfJ3q', 'teat3@mail.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE name ='test1';
DELETE FROM users WHERE name ='test2';
DELETE FROM users WHERE name ='test3';
-- +goose StatementEnd
