CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    surname CHAR(30),
    name CHAR(30),
    patronymic CHAR(30),
    address CHAR(50)
);