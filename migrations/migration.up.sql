CREATE TABLE IF NOT EXISTS app_user(
    id serial PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    age INT
);