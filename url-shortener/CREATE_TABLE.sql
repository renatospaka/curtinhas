CREATE TABLE IF NOT EXISTS urls (
  id serial PRIMARY KEY,
  long_url VARCHAR(150),
  short_url VARCHAR(10)
);
