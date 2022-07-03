CREATE TABLE IF NOT EXISTS users
(
    id            uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    username      varchar(100)     NOT NULL UNIQUE,
    email         varchar(255)     NOT NULL UNIQUE,
    password_hash varchar(255)     NOT NULL,
    created_at    timestamp        NOT NULL DEFAULT (NOW()),
    updated_at    timestamp        NOT NULL DEFAULT (NOW())
);