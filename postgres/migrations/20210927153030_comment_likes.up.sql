CREATE TABLE comment_likes
(
    id        BIGSERIAL PRIMARY KEY,
    commentid   INT NOT NULL,
    useremail VARCHAR(255) NOT NULL
)