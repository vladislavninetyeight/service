CREATE TABLE post (
  post_id BIGSERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES "user"(user_id),
  title varchar(100) NOT NULL,
  body TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);