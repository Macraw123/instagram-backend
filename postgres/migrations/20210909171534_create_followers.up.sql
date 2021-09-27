CREATE TABLE followers
(
    id BIGSERIAL PRIMARY KEY,
    userid BIGSERIAL REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    followerid BIGSERIAL REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL
)
