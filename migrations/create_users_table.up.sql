CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    surname CHAR(30) NOT NULL,
    name CHAR(30) NOT NULL,
    patronymic CHAR(30),
    address CHAR(50) NOT NULL
);