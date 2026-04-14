CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    department TEXT,
    qualification TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);