-- +migrate Up
CREATE TABLE students(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INT NOT NULL,
  enrollment DATE NOT NULL,
  phone VARCHAR(15) NULL
);

-- +migrate Down
DROP TABLE people;
