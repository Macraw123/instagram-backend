CREATE TABLE viewed_stories
(
    id        BIGSERIAL PRIMARY KEY,
    storyid   INT NOT NULL,
    viewed    VARCHAR(255) NOT NULL,
    useremail VARCHAR(255) NOT NULL
)