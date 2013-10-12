-- +goose Up
CREATE TABLE posts (
  id serial,
  title text,
  body text,
  active boolean,
  PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE posts;
