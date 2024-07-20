CREATE TABLE IF NOT EXISTS tasks(
    task_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    name CHAR(50) NOT NULL UNIQUE,
    time INT,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE SET NULL
)