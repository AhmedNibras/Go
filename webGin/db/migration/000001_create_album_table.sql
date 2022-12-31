-- +migrate Up
CREATE TABLE album(
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  artist VARCHAR(255) NOT NULL,
  price NUMERIC(10,2) NOT NULL
);

-- +migrate Down
DROP TABLE album;
