CREATE TABLE search_histories
(
    id        BIGSERIAL PRIMARY KEY,
    word      VARCHAR(255) NOT NULL,
    email     VARCHAR(255) NOT NULL
)