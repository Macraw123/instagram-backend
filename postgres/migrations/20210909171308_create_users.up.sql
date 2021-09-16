CREATE TABLE users
(
    id       BIGSERIAL PRIMARY KEY,
    name     VARCHAR(255)        NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    bio      VARCHAR(255)        NOT NULL,
    email    VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255)        NOT NULL,
    verified BOOLEAN             NOT NULL,
    google   BOOLEAN             NOT NULL,
    token    VARCHAR(255) UNIQUE NOT NULL
)