CREATE TABLE IF NOT EXISTS courses (
                                       id SERIAL PRIMARY KEY,
                                       name VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    duration INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
                             );