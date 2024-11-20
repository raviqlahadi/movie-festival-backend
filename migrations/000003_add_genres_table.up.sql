-- Remove genres column from movies table
ALTER TABLE movies DROP COLUMN IF EXISTS genres;

-- Create genres table with timestamps
CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create movie_genres join table
CREATE TABLE movie_genres (
    movie_id INT NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    genre_id INT NOT NULL REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, genre_id)
);