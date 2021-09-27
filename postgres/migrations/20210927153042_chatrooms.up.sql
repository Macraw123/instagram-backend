CREATE TABLE chatrooms
(
    id        BIGSERIAL PRIMARY KEY,
    owneremail    VARCHAR(255) NOT NULL,
    useremail VARCHAR(255) NOT NULL
)