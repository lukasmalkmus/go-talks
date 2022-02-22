CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL
);

-- name: CreateUser :one
INSERT INTO users (name) VALUES ($1) RETURNING *;