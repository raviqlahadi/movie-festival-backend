-- Add 'year' column to 'movies' table
ALTER TABLE movies
ADD COLUMN year INT NOT NULL DEFAULT 2000;