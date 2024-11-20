-- Drop movie_genres table
DROP TABLE IF EXISTS movie_genres;

-- Drop genres table
DROP TABLE IF EXISTS genres;

-- Add genres column back to movies table
ALTER TABLE movies ADD COLUMN genres TEXT;