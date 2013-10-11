-- +goose Up
CREATE TABLE posts (
  id int NOT NULL,
  title text,
  body text,
  active boolean,
  PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE posts;
