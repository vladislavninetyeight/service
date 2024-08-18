CREATE TABLE "user" (
    user_id BIGSERIAL PRIMARY KEY,
    name varchar(150) NOT NULL,
    phone varchar(12) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);