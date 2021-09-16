CREATE TABLE messages
(
    id       BIGSERIAL PRIMARY KEY,
    toemail    VARCHAR(255) NOT NULL,
    fromemail    VARCHAR(255) NOT NULL,
    created VARCHAR(255) NOT NULL,
    message    VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL
)