CREATE TABLE IF NOT EXISTS link (
  id SERIAL PRIMARY KEY,
  longUrl varchar(255) NOT NULL,
  title varchar(255) NOT NULL,
  createdAt timestamp NOT NULL DEFAULT NOW(),
  shortUrl varchar(255) NOT NULL
);